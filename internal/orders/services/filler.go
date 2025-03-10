package services

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/config"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/evm"
)

type OrderFiller struct {
	ctx      context.Context
	contract *evm.Dexon
	auth     *bind.TransactOpts
	order    *db.Order
}

func newOrderFiller(ctx context.Context, order *db.Order) (*OrderFiller, error) {
	contract, err := evmManager().DexonInstance(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(config.Env.RawPrivKey, txManager().ChainID())
	if err != nil {
		return nil, err
	}

	return &OrderFiller{
		ctx:      ctx,
		contract: contract,
		auth:     auth,
		order:    order,
	}, nil
}

func (f *OrderFiller) executeTransaction(data []byte) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(f.ctx, 2*time.Minute)
	defer cancel()

	return txManager().SendAndWaitForTx(
		ctx,
		f.auth,
		config.Env.DexonContractAddress,
		data,
	)
}

func (f *OrderFiller) handleRejection(orderID int64, err error) (*db.Order, error) {
	log.Error().Err(err).Msg("Failed to execute order")

	rejected := &db.RejectOrderParams{
		ID: orderID,
	}
	_ = rejected.RejectedAt.Scan(time.Now().UTC())

	rejectedOrder, err := db.DB.RejectOrder(f.ctx, *rejected)
	if err != nil {
		return nil, err
	}

	return &rejectedOrder, nil
}

func (f *OrderFiller) createActualQuoteAmount(quoteAmount *big.Int) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   quoteAmount,
		Exp:   -6, // TODO: fix this hardcoded value, this is decimals of USDC
		Valid: true,
	}
}

type orderMapper struct {
	order *db.Order
}

func (m *orderMapper) getBaseFields() (userAddress common.Address, nonce *big.Int, path []byte, amount *big.Int, signature []byte, orderSide uint8, err error) {
	userAddress, err = evm.NormalizeAddress(m.order.Wallet)
	if err != nil {
		return
	}

	nonce = new(big.Int).SetUint64(uint64(m.order.Nonce))

	path, err = evm.NormalizeHex(m.order.Paths)
	if err != nil {
		return
	}

	amount, err = evm.ConvertNumericToDecimals(&m.order.Amount, 0)
	if err != nil {
		err = fmt.Errorf("failed to convert amount to decimals: %w", err)
		return
	}

	signature, err = evm.NormalizeHex(m.order.Signature)
	if err != nil {
		return
	}

	orderSide, err = convertOrderSideToEvmType(m.order.Side)
	if err != nil {
		return
	}

	return
}

func mapOrderToEvmOrder(order *db.Order) (*evm.Order, error) {
	mapper := &orderMapper{order: order}
	userAddress, nonce, path, amount, signature, orderSide, err := mapper.getBaseFields()
	if err != nil {
		return nil, err
	}

	price, err := evm.ConvertDecimalsToWei(&order.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to convert price to wei: %w", err)
	}

	slippage, err := evm.ConvertFloat8ToDecimals(order.Slippage, 6)
	if err != nil {
		return nil, fmt.Errorf("failed to convert slippage to wei: %w", err)
	}

	orderType, err := convertOrderTypeToEvmType(order.Type)
	if err != nil {
		return nil, err
	}

	deadline := new(big.Int).SetInt64(order.Deadline.Time.Unix())

	mapped := &evm.Order{
		Account:      userAddress,
		Nonce:        nonce,
		Path:         path,
		Amount:       amount,
		TriggerPrice: price,
		Slippage:     slippage,
		OrderType:    orderType,
		OrderSide:    orderSide,
		Deadline:     deadline,
		Signature:    signature,
	}

	log.Info().Any("mapped order", mapped).Msg("Mapped order")
	return mapped, nil
}

func mapOrderToEvmTwapOrder(order *db.Order) (*evm.TwapOrder, error) {
	mapper := &orderMapper{order: order}

	userAddress, nonce, path, amount, signature, orderSide, err := mapper.getBaseFields()
	if err != nil {
		return nil, err
	}

	interval := new(big.Int).SetUint64(uint64(order.TwapIntervalSeconds.Int32))

	totalOrders := new(big.Int).SetUint64(uint64(order.TwapCurrentExecutedTimes.Int32))

	mapped := &evm.TwapOrder{
		Account:        userAddress,
		Nonce:          nonce,
		Path:           path,
		Amount:         amount,
		OrderSide:      orderSide,
		Signature:      signature,
		Interval:       interval,
		TotalOrders:    totalOrders,
		StartTimestamp: nil,
	}

	log.Info().Any("mapped TWAP order", mapped).Msg("Mapped TWAP order")
	return mapped, nil
}

func convertOrderTypeToEvmType(orderType db.OrderType) (uint8, error) {
	switch orderType {
	case db.OrderTypeLIMIT:
		return 0, nil
	case db.OrderTypeSTOP:
		return 1, nil
	default:
		return 0, errors.New("invalid order type")
	}
}

func convertOrderSideToEvmType(side db.OrderSide) (uint8, error) {
	switch side {
	case db.OrderSideBUY:
		return 0, nil
	case db.OrderSideSELL:
		return 1, nil
	default:
		return 0, errors.New("invalid order side")
	}
}

func calculateTwapAmount(order *db.Order) *big.Float {
	divisor := big.NewFloat(float64(order.TwapExecutedTimes.Int32))
	f64Amount, _ := order.Amount.Float64Value()
	bigAmount := big.NewFloat(f64Amount.Float64)

	return new(big.Float).Quo(bigAmount, divisor)
}
