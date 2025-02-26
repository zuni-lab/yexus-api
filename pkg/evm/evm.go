package evm

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/config"
)

type Manager struct {
	client      *ethclient.Client
	backoff     *backoff.ExponentialBackOff
	maxAttempts uint64
	handlers    []SwapHandler
}

func NewManager() *Manager {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = 1 * time.Second
	b.MaxInterval = 1 * time.Minute
	b.MaxElapsedTime = 30 * time.Minute

	return &Manager{
		backoff:     b,
		maxAttempts: 100,
	}
}

func (m *Manager) AddHandler(handler SwapHandler) {
	m.handlers = append(m.handlers, handler)
}

func (m *Manager) Close() {
	if m.client != nil {
		m.client.Close()
		m.client = nil
	}
}

func (m *Manager) Connect() error {
	var client *ethclient.Client
	var err error

	operation := func() error {
		log.Info().Msg("Attempting to connect to Ethereum client...")
		client, err = ethclient.Dial(config.Env.AlchemyUrl)
		if err != nil {
			return fmt.Errorf("failed to connect: %w", err)
		}
		return nil
	}

	if err := backoff.Retry(operation, m.backoff); err != nil {
		return err
	}

	m.client = client
	return nil
}

func (m *Manager) WatchPools(ctx context.Context, pools []common.Address) error {
	errChan := make(chan error, len(pools))
	var wg sync.WaitGroup

	for _, pool := range pools {
		wg.Add(1)
		go func(pool common.Address) {
			defer wg.Done()
			if err := m.WatchPool(ctx, pool); err != nil {
				select {
				case errChan <- fmt.Errorf("pool %s: %w", pool.Hex(), err):
				default:
				}
			}
		}(pool)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func (m *Manager) WatchPool(ctx context.Context, pool common.Address) error {
	for {
		if err := m.watchPoolWithRetry(ctx, pool); err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			log.Error().
				Err(err).
				Str("pool", pool.Hex()).
				Msg("Error watching pool, will retry")
			continue
		}
		return nil
	}
}

func (m *Manager) watchPoolWithRetry(ctx context.Context, pool common.Address) error {
	log.Info().Msgf("Watching pool %s", pool.Hex())

	// Determine if we should use WebSocket or polling based on URL
	isWebSocket := strings.HasPrefix(config.Env.AlchemyUrl, "wss")

	if isWebSocket {
		return m.watchPoolWebSocket(ctx, pool)
	}
	return m.watchPoolPolling(ctx, pool)
}

func (m *Manager) watchPoolPolling(ctx context.Context, pool common.Address) error {
	const pollingInterval = 15 * time.Second
	ticker := time.NewTicker(pollingInterval)
	defer ticker.Stop()

	contract, err := NewUniswapV3(pool, m.client)
	if err != nil {
		return fmt.Errorf("failed to create contract instance: %w", err)
	}

	// TODO: remove
	// var lastProcessedBlock = uint64(4_668_400)
	var lastProcessedBlock uint64
	log.Info().Msg("🚀 Ready to watch pool via Polling")

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			currentBlock, err := m.client.BlockNumber(ctx)
			if err != nil {
				log.Error().
					Err(err).
					Str("pool", pool.Hex()).
					Msg("Failed to get current block number")
				continue
			}

			if lastProcessedBlock == 0 {
				lastProcessedBlock = currentBlock - 1
			}

			filterOpts := &bind.FilterOpts{
				Start:   lastProcessedBlock + 1,
				End:     &currentBlock,
				Context: ctx,
			}

			log.Info().Msgf("🐳 [Manager] [watchPoolPolling] with pool %s, start: %d, end: %d", pool.Hex(), lastProcessedBlock+1, currentBlock)

			events, err := contract.FilterSwap(filterOpts, nil, nil)
			if err != nil {
				log.Error().
					Err(err).
					Str("pool", pool.Hex()).
					Msg("Failed to filter swap events")
				continue
			}

			for events.Next() {
				// if err := handler(events.Event); err != nil {
				// 	log.Error().
				// 		Err(err).
				// 		Str("pool", pool.Hex()).
				// 		Msg("Error handling event")
				// }
				for _, handler := range m.handlers {
					if err := handler.HandleSwap(ctx, events.Event); err != nil {
						log.Error().
							Err(err).
							Str("pool", pool.Hex()).
							Msg("Error handling event")
					}
				}
			}

			lastProcessedBlock = currentBlock
		}
	}
}

func (p *Manager) watchPoolWebSocket(ctx context.Context, pool common.Address) error {

	attempt := 0
RETRY:
	for attempt < int(p.maxAttempts) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		contract, err := NewUniswapV3(pool, p.client)
		if err != nil {
			return fmt.Errorf("failed to create contract instance: %w", err)
		}

		watchOpts := &bind.WatchOpts{Context: ctx}
		sink := make(chan *UniswapV3Swap)

		sub, err := contract.WatchSwap(watchOpts, sink, nil, nil)
		if err != nil {
			attempt++
			nextBackoff := p.backoff.NextBackOff()
			if nextBackoff == backoff.Stop {
				return fmt.Errorf("max elapsed time reached after %d attempts", attempt)
			}
			log.Warn().
				Err(err).
				Int("attempt", attempt).
				Dur("backoff", nextBackoff).
				Str("pool", pool.Hex()).
				Msg("Failed to watch swaps, retrying...")
			time.Sleep(nextBackoff)
			continue
		}
		defer sub.Unsubscribe()

		attempt = 0
		p.backoff.Reset()

		log.Info().Msg("🚀 Ready to watch pool via WebSocket")

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case err := <-sub.Err():
				if err != nil {
					log.Warn().
						Err(err).
						Str("pool", pool.Hex()).
						Msg("Subscription error, reconnecting...")
					goto RETRY
				}
			case event := <-sink:
				for _, handler := range p.handlers {
					if err := handler.HandleSwap(ctx, event); err != nil {
						log.Error().
							Err(err).
							Str("pool", pool.Hex()).
							Msg("Error handling event")
					}
				}
			}
		}
	}

	return fmt.Errorf("failed to maintain subscription after %d attempts", p.maxAttempts)
}
