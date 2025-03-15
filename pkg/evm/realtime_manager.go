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
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/yexus-api/config"
	"github.com/zuni-lab/yexus-api/pkg/utils"
)

type RealtimeManager struct {
	*Manager
}

func NewRealtimeManager() *RealtimeManager {
	return &RealtimeManager{
		Manager: NewManager(),
	}
}

func (m *RealtimeManager) WatchPools(ctx context.Context, pools []common.Address) error {
	errChan := make(chan error, len(pools))
	var wg sync.WaitGroup

	for _, pool := range pools {
		wg.Add(1)
		go func(pool common.Address) {
			defer wg.Done()
			if err := m.watchPool(ctx, pool); err != nil {
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

func (m *RealtimeManager) watchPool(ctx context.Context, pool common.Address) error {
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

func (m *RealtimeManager) watchPoolWithRetry(ctx context.Context, pool common.Address) error {
	// Determine if we should use WebSocket or polling based on URL
	isWebSocket := strings.HasPrefix(config.Env.AlchemyUrl, "wss")

	if isWebSocket {
		return m.watchPoolWebSocket(ctx, pool)
	}
	return m.watchPoolPolling(ctx, pool)
}

func (m *RealtimeManager) watchPoolPolling(ctx context.Context, pool common.Address) error {
	contract, err := NewUniswapV3(pool, m.client)
	if err != nil {
		return fmt.Errorf("failed to create contract instance: %w", err)
	}

	ticker := time.NewTicker(config.Env.RealtimeInterval)
	defer ticker.Stop()

	var lastProcessedBlock uint64 = 0

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

			if lastProcessedBlock >= currentBlock-1 {
				log.Debug().
					Uint64("LastProcessedBlock", lastProcessedBlock).
					Uint64("CurrentBlock", currentBlock).
					Msg("No new blocks to process")
				continue
			}

			// Process the new blocks
			startBlock := lastProcessedBlock + 1
			endBlock := currentBlock

			if lastProcessedBlock == 0 {
				startBlock = currentBlock - 1
			}

			log.Info().
				Str("pool", pool.Hex()).
				Uint64("StartBlock", startBlock).
				Uint64("endBlock", endBlock).
				Msg("watchPoolPolling")

			if err := m.processPoolBlockRange(ctx, contract, startBlock, endBlock); err != nil {
				// Check if it's a block availability error
				if strings.Contains(err.Error(), "cannot be found") {
					log.Warn().
						Err(err).
						Str("pool", pool.Hex()).
						Msg("Some blocks not available, adjusting range")
				} else {
					log.Error().
						Err(err).
						Str("pool", pool.Hex()).
						Msg("Error processing block range")
				}
			}

			lastProcessedBlock = currentBlock
		}
	}
}

func (m *RealtimeManager) watchPoolWebSocket(ctx context.Context, pool common.Address) error {

	attempt := 0

	contract, err := NewUniswapV3(pool, m.client)
	if err != nil {
		return fmt.Errorf("failed to create contract instance: %w", err)
	}

RETRY:
	for attempt < int(m.maxAttempts) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		watchOpts := &bind.WatchOpts{Context: ctx}
		sink := make(chan *UniswapV3Swap)

		sub, err := contract.WatchSwap(watchOpts, sink, nil, nil)
		if err != nil {
			attempt++
			nextBackoff := m.backoff.NextBackOff()
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
		m.backoff.Reset()

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
				for _, handler := range m.handlers {
					if err := utils.SafeExecute(ctx, func() error {
						return handler.HandleSwap(ctx, event)
					}); err != nil {
						log.Error().
							Err(err).
							Str("pool", pool.Hex()).
							Msg("Error handling event")
					}
				}
			}
		}
	}

	return fmt.Errorf("failed to maintain subscription after %d attempts", m.maxAttempts)
}
