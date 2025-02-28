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
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/config"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

type RealtimeManager struct {
	*Manager
	pollingInterval time.Duration
}

func NewRealtimeManager() *RealtimeManager {
	return &RealtimeManager{
		Manager:         NewManager(),
		pollingInterval: config.Env.RealtimeInterval,
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

	ticker := time.NewTicker(m.pollingInterval)
	defer ticker.Stop()

	poolAddress := utils.NormalizeAddress(pool.Hex())

	// Get last processed block
	state, err := db.DB.GetBlockProcessingState(ctx, db.GetBlockProcessingStateParams{
		PoolAddress: poolAddress,
		IsBackfill:  false,
	})

	var lastProcessedBlock uint64
	isFirstRun := false

	if err != nil {
		if err != pgx.ErrNoRows {
			return fmt.Errorf("failed to get block processing state: %w", err)
		}

		currentBlock, err := m.client.BlockNumber(ctx)
		if err != nil {
			return fmt.Errorf("failed to get current block number: %w", err)
		}
		lastProcessedBlock = currentBlock - 1
		isFirstRun = true // Mark the first run

		// initialize the state
		if err := db.DB.UpsertBlockProcessingState(ctx, db.UpsertBlockProcessingStateParams{
			PoolAddress:        poolAddress,
			IsBackfill:         false,
			LastProcessedBlock: int64(lastProcessedBlock),
		}); err != nil {
			return fmt.Errorf("failed to initialize block processing state: %w", err)
		}
	} else {
		lastProcessedBlock = uint64(state.LastProcessedBlock)
	}

	log.Info().
		Uint64("last_processed_block", lastProcessedBlock).
		Bool("first_run", isFirstRun).
		Str("pool", pool.Hex()).
		Msg("🚀 Ready to watch pool via Polling")

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

			// Calculate how many new blocks are available
			newBlocks := currentBlock - lastProcessedBlock

			// Skip if no new blocks
			if newBlocks == 0 {
				log.Debug().
					Str("pool", pool.Hex()).
					Msg("No new blocks to process")
				continue
			}

			// Check if we have enough blocks to process
			if newBlocks < config.Env.RealtimeMinBlockRange {
				// Special cases to process anyway:
				// 1. First polling cycle after initialization
				// 2. When we have a significant gap (more than 5 blocks)
				if !isFirstRun {
					log.Debug().
						Uint64("blocks", newBlocks).
						Uint64("minimum", config.Env.RealtimeMinBlockRange).
						Str("pool", pool.Hex()).
						Msg("Not enough blocks to process yet, waiting for more")
					continue
				}

				log.Info().
					Uint64("blocks", newBlocks).
					Bool("first_run", isFirstRun).
					Str("pool", pool.Hex()).
					Msg("Processing blocks despite being below threshold")
			} else if newBlocks > 100 {
				log.Warn().
					Uint64("blocks", newBlocks).
					Str("pool", pool.Hex()).
					Msg("Processing unusually large block range")
			}

			// Process the new blocks
			startBlock := lastProcessedBlock + 1
			endBlock := currentBlock

			log.Info().
				Uint64("start", startBlock).
				Uint64("end", endBlock).
				Uint64("count", endBlock-startBlock+1).
				Str("pool", pool.Hex()).
				Msg("Processing block range")

			if err := m.processPoolBlockRange(ctx, contract, startBlock, endBlock); err != nil {
				// Check if it's a block availability error
				if strings.Contains(err.Error(), "cannot be found") {
					log.Warn().
						Err(err).
						Str("pool", pool.Hex()).
						Msg("Some blocks not available, adjusting range")

					// Skip ahead to avoid getting stuck
					if err := db.DB.UpsertBlockProcessingState(ctx, db.UpsertBlockProcessingStateParams{
						PoolAddress:        poolAddress,
						IsBackfill:         false,
						LastProcessedBlock: int64(currentBlock),
					}); err != nil {
						log.Error().Err(err).Msg("Failed to update block processing state")
					}

					lastProcessedBlock = currentBlock
				} else {
					log.Error().
						Err(err).
						Str("pool", pool.Hex()).
						Msg("Error processing block range")
				}
				continue
			}

			if err := db.DB.UpsertBlockProcessingState(ctx, db.UpsertBlockProcessingStateParams{
				PoolAddress:        poolAddress,
				IsBackfill:         false,
				LastProcessedBlock: int64(currentBlock),
			}); err != nil {
				log.Error().
					Err(err).
					Msg("Failed to update block processing state")
				continue
			}

			lastProcessedBlock = currentBlock
			isFirstRun = false
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
