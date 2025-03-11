package evm

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

type TxManager struct {
	client          *ethclient.Client
	chainID         *big.Int
	nonceCache      sync.Map // thread-safe map for nonce caching
	gasMultiplier   float64
	maxGasPrice     *big.Int
	maxNonceRetries int
}

func NewTxManager(client *ethclient.Client) (*TxManager, error) {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	return &TxManager{
		client:          client,
		chainID:         chainID,
		gasMultiplier:   1.1,               // 10% buffer for gas estimation
		maxGasPrice:     big.NewInt(500e9), // 500 gwei max
		maxNonceRetries: 5,
	}, nil
}

func (tm *TxManager) ChainID() *big.Int {
	return tm.chainID
}

func (m *TxManager) SendAndWaitForTxWithNonceRetry(ctx context.Context, auth *bind.TransactOpts, to common.Address, data []byte) (*types.Receipt, error) {
	var (
		receipt *types.Receipt
		err     error
	)

	for attempt := 0; attempt < m.maxNonceRetries; attempt++ {
		if attempt > 0 {
			if err := m.refreshNonce(ctx, auth.From); err != nil {
				log.Error().Err(err).Msg("failed to refresh nonce")
				continue
			}
		}

		receipt, err = m.sendAndWaitForTx(ctx, auth, to, data)
		if err == nil {
			return receipt, nil
		}

		if !strings.Contains(err.Error(), "nonce too low") {
			return nil, err
		}

		log.Warn().
			Err(err).
			Int("attempt", attempt+1).
			Msg("nonce too low, retrying with new nonce")
	}

	return nil, fmt.Errorf("failed after %d nonce retry attempts: %w", m.maxNonceRetries, err)
}

func (tm *TxManager) sendAndWaitForTx(ctx context.Context, opts *bind.TransactOpts, to common.Address, data []byte) (*types.Receipt, error) {
	var err error

	nonce, err := tm.getNonce(ctx, opts.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	gasLimit, err := tm.estimateGas(ctx, opts, to, data)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// 3. Get current gas price with safety checks
	gasPrice, err := tm.getSafeGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	// 4. Create transaction
	tx := types.NewTransaction(
		nonce,
		to,
		opts.Value,
		gasLimit,
		gasPrice,
		data,
	)

	// 5. Sign transaction
	signedTx, err := opts.Signer(opts.From, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to sign tx: %w", err)
	}

	// 6. Send transaction with retry logic
	err = tm.sendTxWithRetry(ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send tx: %w", err)
	}

	// 7. Wait for transaction with timeout
	receipt, err := tm.waitForTx(ctx, signedTx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed waiting for tx: %w", err)
	}

	log.Info().
		Str("txHash", receipt.TxHash.String()).
		Uint64("gasUsed", receipt.GasUsed).
		Msg("🎯 Transaction successful")

	// 8. Update nonce cache on success
	tm.nonceCache.Store(opts.From, nonce+1)

	return receipt, nil
}

func (m *TxManager) refreshNonce(ctx context.Context, address common.Address) error {
	nonce, err := m.client.PendingNonceAt(ctx, address)
	if err != nil {
		return fmt.Errorf("failed to refresh nonce: %w", err)
	}

	m.nonceCache.Store(address, nonce)
	return nil
}

func (tm *TxManager) getNonce(ctx context.Context, from common.Address) (uint64, error) {
	// Check cache first
	if cachedNonce, ok := tm.nonceCache.Load(from); ok {
		return cachedNonce.(uint64), nil
	}

	// Get pending nonce from chain
	nonce, err := tm.client.PendingNonceAt(ctx, from)
	if err != nil {
		return 0, err
	}

	tm.nonceCache.Store(from, nonce)
	return nonce, nil
}

func (tm *TxManager) estimateGas(ctx context.Context, opts *bind.TransactOpts, to common.Address, data []byte) (uint64, error) {
	msg := ethereum.CallMsg{
		From:     opts.From,
		To:       &to,
		Gas:      0,
		GasPrice: opts.GasPrice,
		Value:    opts.Value,
		Data:     data,
	}

	estimatedGas, err := tm.client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, err
	}

	// Add buffer to estimated gas
	gasLimit := uint64(float64(estimatedGas) * tm.gasMultiplier)
	return gasLimit, nil
}

func (tm *TxManager) getSafeGasPrice(ctx context.Context) (*big.Int, error) {
	gasPrice, err := tm.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	// Ensure gas price doesn't exceed maximum
	if gasPrice.Cmp(tm.maxGasPrice) > 0 {
		return tm.maxGasPrice, nil
	}

	return gasPrice, nil
}

func (tm *TxManager) sendTxWithRetry(ctx context.Context, tx *types.Transaction) error {
	var err error
	for i := 0; i < 3; i++ {
		err = tm.client.SendTransaction(ctx, tx)
		if err == nil {
			return nil
		}

		// If nonce too low, break immediately
		if strings.Contains(err.Error(), "nonce too low") {
			return err
		}

		// Wait before retry
		time.Sleep(time.Second * time.Duration(i+1))
	}
	return err
}

func (tm *TxManager) waitForTx(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	log.Info().Str("txHash", txHash.String()).Msg("🌐 Waiting for transaction")

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			receipt, err := tm.client.TransactionReceipt(ctx, txHash)
			if err != nil {
				if err == ethereum.NotFound {
					continue
				}
				return nil, err
			}

			// Check if transaction was successful
			if receipt.Status == types.ReceiptStatusSuccessful {
				return receipt, nil
			}

			// Get transaction failure reason
			reason, err := tm.getRevertReason(ctx, txHash)
			if err != nil {
				return nil, fmt.Errorf("transaction failed: %w", err)
			}
			return nil, fmt.Errorf("transaction reverted: %s", reason)
		}
	}
}

func (tm *TxManager) getRevertReason(ctx context.Context, txHash common.Hash) (string, error) {
	tx, _, err := tm.client.TransactionByHash(ctx, txHash)
	if err != nil {
		return "", err
	}

	from, err := types.Sender(types.NewEIP155Signer(tm.chainID), tx)
	if err != nil {
		return "", err
	}

	// Call the transaction to get the revert reason
	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	result, err := tm.client.CallContract(ctx, msg, nil)
	if err != nil {
		return err.Error(), nil
	}

	return hex.EncodeToString(result), nil
}
