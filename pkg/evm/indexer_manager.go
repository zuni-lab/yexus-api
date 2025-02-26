package evm

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/config"
)

type IndexerManager struct {
	*Manager
	chunkSize   uint64
	concurrency int
}

func NewIndexerManager() *IndexerManager {
	return &IndexerManager{
		Manager:     NewManager(),
		chunkSize:   config.Env.ChunkSize,
		concurrency: config.Env.Concurrency,
	}
}

func (m *IndexerManager) SetChunkSize(size uint64) {
	m.chunkSize = size
}

func (m *IndexerManager) SetConcurrency(n int) {
	m.concurrency = n
}

func (m *IndexerManager) IndexPools(ctx context.Context, pools []common.Address, fromBlock uint64) error {
	currentBlock, err := m.client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get current block: %w", err)
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, m.concurrency)
	errChan := make(chan error, 1)

	for _, pool := range pools {

		contract, err := NewUniswapV3(pool, m.client)
		if err != nil {
			return fmt.Errorf("failed to create contract instance: %w", err)
		}

		startBlock := fromBlock
		for startBlock <= currentBlock {
			endBlock := startBlock + m.chunkSize - 1
			if endBlock > currentBlock {
				endBlock = currentBlock
			}

			wg.Add(1)
			sem <- struct{}{} // semaphore

			go func(pool common.Address, start, end uint64) {
				defer func() {
					<-sem // release semaphore
					wg.Done()
				}()

				log.Info().
					Str("pool", pool.Hex()).
					Uint64("start", start).
					Uint64("end", end).
					Msg("Processing block range")

				if err := m.processPoolBlockRange(ctx, contract, start, end); err != nil {
					select {
					case errChan <- fmt.Errorf("failed to process pool %s blocks %d-%d: %w",
						pool.Hex(), start, end, err):
					default:
					}
				}
			}(pool, startBlock, endBlock)

			startBlock = endBlock + 1
		}
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Return first error if any
	if err := <-errChan; err != nil {
		return err
	}

	return nil
}
