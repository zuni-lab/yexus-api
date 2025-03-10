package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/config"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/evm"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

var (
	evmManager = sync.OnceValue(func() *evm.Manager {
		log.Info().Msg("Connecting to EVM manager")
		manager := evm.NewManager()
		if err := manager.Connect(); err != nil {
			panic(err)
		}
		return manager
	})

	txManager = sync.OnceValue(func() *evm.TxManager {
		log.Info().Msg("Creating transaction manager")
		manager, err := evm.NewTxManager(evmManager().Client())
		if err != nil {
			panic(err)
		}
		return manager
	})
)

func MatchOrder(ctx context.Context, price *big.Float) (*db.Order, error) {
	numericPrice, err := utils.BigFloatToNumeric(price)
	if err != nil {
		return nil, err
	}

	order, err := db.DB.GetMatchedOrder(ctx, numericPrice)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("no order matched")
		}
		return nil, err
	}

	log.Info().Any("matched orders", order).Msg("Matched order")

	var filledOrder *db.Order
	if order.Type == db.OrderTypeTWAP {
		filledOrder, err = fillTwapOrder(ctx, &order, price)
	} else {
		filledOrder, err = fillOrder(ctx, &order)
	}

	if err != nil {
		return nil, err
	}

	return filledOrder, nil
}

func MatchTwapOrders() {
	orders, err := db.DB.GetMatchedTwapOrder(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
			log.Warn().Err(err).Msg("⚠️ [SwapHandler] No matched TWAP orders")
		}
		log.Warn().Err(err).Msg("⚠️ [SwapHandler] Failed to get matched TWAP orders")
	}

	fmt.Println("hello", orders)

	for _, order := range orders {
		_, err = fillTwapOrder(context.Background(), &order, new(big.Float).SetFloat64(0))
		if err != nil {
			log.Warn().Any("id", order.ID).Err(err).Msg("⚠️ [SwapHandler] Failed to match TWAP order")
		} else {
			log.Info().Any("id", order.ID).Msg("✅ [SwapHandler] Successfully matched TWAP order")
		}
	}
}

func fillOrder(ctx context.Context, order *db.Order) (*db.Order, error) {
	contract, err := evmManager().DexonInstance(ctx)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	auth, err := bind.NewKeyedTransactorWithChainID(config.Env.RawPrivKey, txManager().ChainID())
	if err != nil {
		return nil, err
	}

	mappedOrder, err := mapOrderToEvmOrder(order)
	if err != nil {
		return nil, err
	}

	data, err := evm.ExecuteOrderData(&contract.DexonTransactor, mappedOrder)
	if err != nil {
		return nil, err
	}

	receipt, err := txManager().SendAndWaitForTx(
		ctx,
		auth,
		config.Env.DexonContractAddress,
		data,
	)

	var rejected *db.RejectOrderParams

	if err != nil {
		log.Error().Err(err).Msg("Failed to send and wait for transaction")

		rejected = &db.RejectOrderParams{
			ID: order.ID,
		}

		_ = rejected.RejectedAt.Scan(time.Now().UTC())
	}

	event, err := evm.ParseOrderExecutedEvent(&contract.DexonFilterer, receipt)
	if err != nil {
		rejected = &db.RejectOrderParams{
			ID: order.ID,
		}
		_ = rejected.RejectedAt.Scan(time.Now().UTC())
	}

	if rejected != nil {
		rejectedOrder, err := db.DB.RejectOrder(ctx, *rejected)
		if err != nil {
			return nil, err
		}

		return &rejectedOrder, nil
	}

	actualAmount := pgtype.Numeric{
		Int:   event.ActualSwapAmount,
		Exp:   -6, // TODO: fix this hardcoded value, this is decimals of USDC
		Valid: true,
	}

	params := db.FillOrderParams{
		ID:           order.ID,
		ActualAmount: actualAmount,
	}
	_ = params.FilledAt.Scan(time.Now().UTC())
	_ = params.TxHash.Scan(receipt.TxHash.String())

	filledOrder, err := db.DB.FillOrder(ctx, params)
	if err != nil {
		return nil, err
	}

	return &filledOrder, nil
}

func mapOrderToEvmOrder(order *db.Order) (*evm.Order, error) {
	userAddress, err := evm.NormalizeAddress(order.Wallet)
	if err != nil {
		return nil, err
	}

	nonce := new(big.Int).SetUint64(uint64(order.Nonce))

	path, err := evm.NormalizeHex(order.Paths)
	if err != nil {
		return nil, err
	}

	// Not need to convert to wei because the input of the client is already in wei
	amount, err := evm.ConvertNumericToDecimals(&order.Amount, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to convert amount to decimals: %w", err)
	}

	price, err := evm.ConvertDecimalsToWei(&order.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to convert price to wei: %w", err)
	}

	slippage, err := evm.ConvertFloat8ToDecimals(order.Slippage, 6)
	if err != nil {
		return nil, fmt.Errorf("failed to convert slippage to wei: %w", err)
	}

	deadline := new(big.Int).SetInt64(order.Deadline.Time.Unix())

	signature, err := evm.NormalizeHex(order.Signature)
	if err != nil {
		return nil, err
	}

	orderType, err := convertOrderTypeToEvmType(order.Type)
	if err != nil {
		return nil, err
	}

	orderSide, err := convertOrderSideToEvmType(order.Side)
	if err != nil {
		return nil, err
	}

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

func fillTwapOrder(ctx context.Context, order *db.Order, price *big.Float) (*db.Order, error) {
	var (
		params = db.FillTwapOrderParams{
			ID: order.ID,
		}
		now = time.Now().UTC()
		err error
	)

	_ = params.FilledAt.Scan(now)
	_ = params.TwapCurrentExecutedTimes.Scan(order.TwapCurrentExecutedTimes.Int32 + 1)
	if order.TwapCurrentExecutedTimes.Int32+1 == order.TwapExecutedTimes.Int32 {
		params.Status = db.OrderStatusFILLED
		_ = params.FilledAt.Scan(now)
	} else {
		params.Status = db.OrderStatusPARTIALFILLED
		_ = params.PartialFilledAt.Scan(now)
	}

	amount := calculateTwapAmount(order)
	err = fillPartialOrder(ctx, order, price, amount, now)
	if err != nil {
		return nil, err
	}

	filledOrder, err := db.DB.FillTwapOrder(ctx, params)
	if err != nil {
		return nil, err
	}

	return &filledOrder, nil
}

func calculateTwapAmount(order *db.Order) *big.Float {
	divisor := big.NewFloat(float64(order.TwapExecutedTimes.Int32))
	f64Amount, _ := order.Amount.Float64Value()
	bigAmount := big.NewFloat(f64Amount.Float64)

	return new(big.Float).Quo(bigAmount, divisor)
}

func fillPartialOrder(ctx context.Context, parent *db.Order, price, amount *big.Float, now time.Time) error {
	params := db.InsertOrderParams{
		PoolIds: parent.PoolIds,
		Wallet:  parent.Wallet,
		Status:  db.OrderStatusFILLED,
		Side:    parent.Side,
		Type:    db.OrderTypeTWAP,
		Paths:   parent.Paths,
	}

	_ = params.ParentID.Scan(parent.ID)
	_ = params.Price.Scan(price.String())
	_ = params.Amount.Scan(amount.String())
	_ = params.FilledAt.Scan(now)
	params.CreatedAt = params.FilledAt

	_, err := db.DB.InsertOrder(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
