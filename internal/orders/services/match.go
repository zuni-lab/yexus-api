package services

import (
	"context"
	"database/sql"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
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
			return nil, errors.New("no new orders to match")
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
		log.Warn().Err(err).Msg("⚠️ [SwapHandler] failed to get matched TWAP orders")
		return
	}

	if len(orders) == 0 {
		log.Warn().Err(err).Msg("⚠️ [SwapHandler] no new TWAP orders to match")
		return
	}

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
	filler, err := newOrderFiller(ctx, order)
	if err != nil {
		return nil, err
	}

	mappedOrder, err := mapOrderToEvmOrder(order)
	if err != nil {
		return nil, err
	}

	data, err := evm.ExecuteOrderData(&filler.contract.DexonTransactor, mappedOrder)
	if err != nil {
		return nil, err
	}

	receipt, err := filler.executeTransaction(data)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send and wait for transaction")
		return filler.handleRejection(order.ID)
	}

	event, err := evm.ParseOrderExecutedEvent(&filler.contract.DexonFilterer, receipt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse order executed event")
		return filler.handleRejection(order.ID)
	}

	params := db.FillOrderParams{
		ID:           order.ID,
		ActualAmount: filler.createActualQuoteAmount(event.QuoteAmount),
	}
	_ = params.FilledAt.Scan(time.Now().UTC())
	_ = params.TxHash.Scan(receipt.TxHash.String())

	filledOrder, err := db.DB.FillOrder(ctx, params)
	if err != nil {
		return nil, err
	}

	return &filledOrder, nil
}

func fillTwapOrder(ctx context.Context, order *db.Order, price *big.Float) (*db.Order, error) {
	now := time.Now().UTC()

	params := createTwapFillParams(order, now)
	amount := calculateTwapAmount(order)

	var err error
	params.ActualAmount, err = fillPartialOrder(ctx, order, price, amount, now)
	if err != nil {
		return nil, err
	}

	filledOrder, err := db.DB.FillTwapOrder(ctx, params)
	if err != nil {
		return nil, err
	}

	return &filledOrder, nil
}

func fillPartialOrder(ctx context.Context, parent *db.Order, price, amount *big.Float, now time.Time) (pgtype.Numeric, error) {
	var actualAmount pgtype.Numeric

	filler, err := newOrderFiller(ctx, parent)
	if err != nil {
		return actualAmount, err
	}

	mappedOrder, err := mapOrderToEvmTwapOrder(parent)
	if err != nil {
		return actualAmount, err
	}

	data, err := evm.ExecuteTwapOrderData(&filler.contract.DexonTransactor, mappedOrder)
	if err != nil {
		return actualAmount, err
	}

	receipt, err := filler.executeTransaction(data)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send and wait for twap transaction")
		_, rejectionErr := filler.handleRejection(parent.ID)
		if rejectionErr != nil {
			log.Error().Err(rejectionErr).Msg("Failed to reject twap order")
		}
		return actualAmount, err
	}

	event, err := evm.ParseTwapOrderExecutedEvent(&filler.contract.DexonFilterer, receipt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse twap order executed event")
		_, rejectionErr := filler.handleRejection(parent.ID)
		if rejectionErr != nil {
			log.Error().Err(rejectionErr).Msg("Failed to reject twap order")
		}
		return actualAmount, err
	}

	actualAmount = filler.createActualQuoteAmount(event.QuoteAmount)
	params := createPartialOrderParams(parent, price, amount, receipt.TxHash.String(), actualAmount, now)

	_, err = db.DB.InsertOrder(ctx, params)
	return actualAmount, err
}

func createTwapFillParams(order *db.Order, now time.Time) db.FillTwapOrderParams {
	params := db.FillTwapOrderParams{
		ID: order.ID,
	}

	nextExecutionCount := order.TwapCurrentExecutedTimes.Int32 + 1
	_ = params.TwapCurrentExecutedTimes.Scan(int64(nextExecutionCount))

	if nextExecutionCount == order.TwapExecutedTimes.Int32 {
		params.Status = db.OrderStatusFILLED
		_ = params.FilledAt.Scan(now)
	} else {
		params.Status = db.OrderStatusPARTIALFILLED
		_ = params.PartialFilledAt.Scan(now)
	}

	return params
}

func createPartialOrderParams(parent *db.Order, price, amount *big.Float, txHash string,
	actualAmount pgtype.Numeric, now time.Time) db.InsertOrderParams {

	params := db.InsertOrderParams{
		PoolIds:      parent.PoolIds,
		Wallet:       parent.Wallet,
		Status:       db.OrderStatusFILLED,
		Side:         parent.Side,
		Type:         db.OrderTypeTWAP,
		Paths:        parent.Paths,
		Signature:    parent.Signature,
		Nonce:        parent.Nonce,
		ActualAmount: actualAmount,
	}

	_ = params.ParentID.Scan(parent.ID)
	_ = params.Price.Scan(price.String())
	_ = params.TxHash.Scan(txHash)
	_ = params.FilledAt.Scan(now)
	params.Amount, _ = utils.BigFloatToNumeric(amount)
	params.CreatedAt = params.FilledAt

	return params
}
