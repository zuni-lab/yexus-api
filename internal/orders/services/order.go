package services

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

type ListOrdersByWalletQuery struct {
	Wallet    string           `query:"wallet" validate:"eth_addr"`
	Status    []db.OrderStatus `query:"status" validate:"dive,oneof=PENDING PARTIAL_FILLED FILLED REJECTED CANCELLED"`
	NotStatus []db.OrderStatus `query:"not_status" validate:"dive,oneof=PENDING PARTIAL_FILLED FILLED REJECTED CANCELLED"`
	Types     []db.OrderType   `query:"types" validate:"dive,oneof=MARKET LIMIT STOP TWAP"`
	Side      *string          `query:"side" validate:"omitempty,oneof=BUY SELL"`
	ParentID  *int64           `query:"parentId" validate:"omitempty,gt=0"`
	Limit     int32            `query:"limit" validate:"gt=0"`
	Offset    int32            `query:"offset" validate:"gte=0"`
}

func ListOrderByWallet(ctx context.Context, query ListOrdersByWalletQuery) (*db.ListOrdersByWalletResult, error) {
	query.Wallet = utils.NormalizeAddress(query.Wallet)

	var params db.GetOrdersByWalletParams
	if err := copier.Copy(&params, &query); err != nil {
		return nil, err
	}

	result, err := db.DB.ListOrdersByWalletTx(ctx, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type GetOrderByIDQuery struct {
	ID     int64
	Wallet string `query:"wallet" validate:"eth_addr"`
}

func GetOrderByID(ctx context.Context, query GetOrderByIDQuery) (*db.GetOrderByIDRow, error) {
	query.Wallet = utils.NormalizeAddress(query.Wallet)

	var params db.GetOrderByIDParams
	if err := copier.Copy(&params, &query); err != nil {
		return nil, err
	}

	order, err := db.DB.GetOrderByID(ctx, params)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

type CreateOrderBody struct {
	Wallet    string       `json:"wallet" validate:"eth_addr"`
	PoolIDs   []string     `json:"poolIds" validate:"min=1,dive,eth_addr"`
	Side      db.OrderSide `json:"side" validate:"oneof=BUY SELL"`
	Type      db.OrderType `json:"type" validate:"oneof=MARKET LIMIT STOP TWAP"`
	Price     *string      `json:"price" validate:"required_unless=Type TWAP,numeric,gt=0"`
	Amount    string       `json:"amount" validate:"numeric,gt=0"`
	Slippage  float64      `json:"slippage" validate:"gte=0"`
	Signature string       `json:"signature"`
	Paths     string       `json:"paths" validate:"max=256"`
	Nonce     string       `json:"nonce"`
	Deadline  *int64       `json:"deadline" validate:"omitempty,gt=0"`

	TwapIntervalSeconds *int64  `json:"twapIntervalSeconds" validate:"required_if=Type TWAP,omitempty,gt=59"`
	TwapExecutedTimes   *int64  `json:"twapExecutedTimes" validate:"required_if=Type TWAP,omitempty,gt=0"`
	TwapMinPrice        *string `json:"twapMinPrice" validate:"omitempty,numeric,gte=0"`
	TwapMaxPrice        *string `json:"twapMaxPrice" validate:"required_with=TwapMinPrice,omitempty,numeric,gtefield=TwapMinPrice"`
	TwapStartedAt       *int64  `json:"twapStartedAt" validate:"required_if=Type TWAP,omitempty,gt=0"`
}

func CreateOrder(ctx context.Context, body CreateOrderBody) (*db.InsertOrderRow, error) {
	body.Wallet = utils.NormalizeAddress(body.Wallet)

	var (
		params db.InsertOrderParams
		err    error
	)
	if err := copier.Copy(&params, &body); err != nil {
		return nil, err
	}

	params.Nonce, err = strconv.ParseInt(body.Nonce, 10, 64)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	_ = params.CreatedAt.Scan(now)
	if params.Type == db.OrderTypeMARKET {
		_ = params.FilledAt.Scan(now)
		params.CreatedAt = params.FilledAt
		params.Status = db.OrderStatusFILLED
	} else {
		params.Status = db.OrderStatusPENDING

		if params.Type == db.OrderTypeTWAP {
			_ = params.Price.Scan("0")
			_ = params.Slippage.Scan(nil)
			_ = params.TwapCurrentExecutedTimes.Scan(int64(0))

			if body.TwapStartedAt != nil {
				_ = params.TwapStartedAt.Scan(time.Unix(*body.TwapStartedAt, 0).UTC())
			}
		} else {
			params.TwapIntervalSeconds.Valid = false
			params.TwapExecutedTimes.Valid = false
			params.TwapMinPrice.Valid = false
			params.TwapMaxPrice.Valid = false
			params.TwapStartedAt.Valid = false
			_ = params.Deadline.Scan(time.Unix(*body.Deadline, 0).UTC())
		}
	}

	if body.Deadline != nil && body.Type != db.OrderTypeTWAP {
		if now.Unix() >= *body.Deadline {
			return nil, errors.New("invalid deadline")
		}

		_ = params.Deadline.Scan(time.Unix(*body.Deadline, 0).UTC())
	}

	for i, poolID := range body.PoolIDs {
		body.PoolIDs[i] = utils.NormalizeAddress(poolID)
	}

	pools, err := db.DB.GetPoolsByIDs(ctx, body.PoolIDs)
	if err != nil {
		return nil, err
	} else if len(pools) != len(body.PoolIDs) {
		return nil, errors.New("invalid pool ids")
	}

	order, err := db.DB.InsertOrder(ctx, params)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

type CancelOrderBody struct {
	ID     int64
	Wallet string `json:"wallet" validate:"eth_addr"`
}

func CancelOrder(ctx context.Context, body CancelOrderBody) (*db.CancelOrderRow, error) {
	body.Wallet = utils.NormalizeAddress(body.Wallet)

	var params db.CancelOrderParams
	if err := copier.Copy(&params, &body); err != nil {
		return nil, err
	}

	_ = params.CancelledAt.Scan(time.Now().UTC())
	order, err := db.DB.CancelOrder(ctx, params)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

type CancelAllOrdersBody struct {
	Wallet string `json:"wallet" validate:"eth_addr"`
}

func CancelAllOrders(ctx context.Context, body CancelAllOrdersBody) error {
	body.Wallet = utils.NormalizeAddress(body.Wallet)

	var params db.CancelAllOrdersParams
	if err := copier.Copy(&params, &body); err != nil {
		return err
	}

	_ = params.CancelledAt.Scan(time.Now().UTC())
	err := db.DB.CancelAllOrders(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
