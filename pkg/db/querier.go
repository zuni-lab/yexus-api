// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CancelAllOrders(ctx context.Context, arg CancelAllOrdersParams) error
	CancelOrder(ctx context.Context, arg CancelOrderParams) (CancelOrderRow, error)
	CountChatThreads(ctx context.Context, userAddress string) (int64, error)
	CountOrdersByWallet(ctx context.Context, arg CountOrdersByWalletParams) (int64, error)
	CreatePool(ctx context.Context, arg CreatePoolParams) (Pool, error)
	CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error)
	FillOrder(ctx context.Context, arg FillOrderParams) (Order, error)
	FillTwapOrder(ctx context.Context, arg FillTwapOrderParams) (Order, error)
	GetChatThread(ctx context.Context, arg GetChatThreadParams) (ChatThread, error)
	GetChatThreads(ctx context.Context, arg GetChatThreadsParams) ([]ChatThread, error)
	GetLatestYieldMetric(ctx context.Context) (YieldMetric, error)
	GetMatchedOrder(ctx context.Context, price pgtype.Numeric) (Order, error)
	GetMatchedTwapOrder(ctx context.Context) ([]Order, error)
	GetOrderByID(ctx context.Context, arg GetOrderByIDParams) (GetOrderByIDRow, error)
	GetOrdersByWallet(ctx context.Context, arg GetOrdersByWalletParams) ([]GetOrdersByWalletRow, error)
	GetPool(ctx context.Context, id string) (Pool, error)
	GetPoolByToken(ctx context.Context, arg GetPoolByTokenParams) (Pool, error)
	GetPools(ctx context.Context) ([]Pool, error)
	GetPoolsByIDs(ctx context.Context, ids []string) ([]Pool, error)
	GetYieldMetrics(ctx context.Context, arg GetYieldMetricsParams) ([]YieldMetric, error)
	GetYieldMetricsForChat(ctx context.Context, dollar_1 []string) ([]GetYieldMetricsForChatRow, error)
	InsertOrder(ctx context.Context, arg InsertOrderParams) (InsertOrderRow, error)
	PoolDetails(ctx context.Context, id string) (PoolDetailsRow, error)
	RejectOrder(ctx context.Context, arg RejectOrderParams) (Order, error)
	UpsertChatThread(ctx context.Context, arg UpsertChatThreadParams) (ChatThread, error)
}

var _ Querier = (*Queries)(nil)
