package evm

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/config"
)

type Manager struct {
	client        *ethclient.Client
	backoff       *backoff.ExponentialBackOff
	maxAttempts   uint64
	handlers      []SwapHandler
	chainID       *big.Int
	dexonContract *Dexon
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

func (m *Manager) Connect() error {
	var (
		client *ethclient.Client
		err    error
	)

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

func (m *Manager) Close() {
	if m.client != nil {
		m.client.Close()
		m.client = nil
	}
}

// core function to process a range of blocks for a given pool
func (m *Manager) processPoolBlockRange(ctx context.Context, contract *UniswapV3, start, end uint64) error {
	filterOpts := &bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	events, err := contract.FilterSwap(filterOpts, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to filter swap events: %w", err)
	}

	// TODO: Get latest event from the contract

	var i uint64

	for events.Next() {
		for _, handler := range m.handlers {
			if err := handler.HandleSwap(ctx, events.Event); err != nil {
				log.Error().
					Err(err).
					Msg("Error handling event")
			}
		}
		i += 1
	}

	log.Info().
		Uint64("Number of events", i).
		Msg("processPoolBlockRange")

	return nil
}

func (m *Manager) Client() *ethclient.Client {
	return m.client
}

func (m *Manager) DexonInstance(ctx context.Context) (*Dexon, error) {
	if m.dexonContract != nil {
		return m.dexonContract, nil
	}
	contractAddress := common.HexToAddress(config.Env.ContractAddress)
	return NewDexon(contractAddress, m.client)
}

func (m *Manager) ChainID(ctx context.Context) (*big.Int, error) {
	if m.chainID != nil {
		return m.chainID, nil
	}

	chainID, err := m.Client().NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	m.chainID = chainID
	return chainID, nil
}
