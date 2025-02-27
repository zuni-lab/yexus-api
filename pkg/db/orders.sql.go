// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: orders.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getMatchedOrder = `-- name: GetMatchedOrder :one
SELECT id, pool_id, parent_id, wallet, status, side, type, price, amount, twap_amount, twap_parts, partial_filled_at, filled_at, cancelled_at, created_at FROM orders
WHERE (
        (side = 'BUY' AND type = 'LIMIT' AND price <= $1)
        OR (side = 'SELL' AND type = 'LIMIT' AND price >= $1)
        OR (side = 'BUY' AND type = 'STOP' AND price >= $1)
        OR (side = 'SELL' AND type 'STOP' AND price <= $1)
        OR (side = 'BUY' AND type = 'TWAP' AND price <= $1)
        OR (side = 'SELL' AND type = 'TWAP' AND price >= $1)
    )
    AND status IN ('PENDING', 'PARTIAL_FILLED')
    AND type <> 'TWAP'
ORDER BY created_at ASC
LIMIT 1
`

func (q *Queries) GetMatchedOrder(ctx context.Context, price pgtype.Numeric) (Order, error) {
	row := q.db.QueryRow(ctx, getMatchedOrder, price)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.PoolID,
		&i.ParentID,
		&i.Wallet,
		&i.Status,
		&i.Side,
		&i.Type,
		&i.Price,
		&i.Amount,
		&i.TwapAmount,
		&i.TwapParts,
		&i.PartialFilledAt,
		&i.FilledAt,
		&i.CancelledAt,
		&i.CreatedAt,
	)
	return i, err
}

const getOrdersByStatus = `-- name: GetOrdersByStatus :many
SELECT id, pool_id, parent_id, wallet, status, side, type, price, amount, twap_amount, twap_parts, partial_filled_at, filled_at, cancelled_at, created_at FROM orders
WHERE status = ANY($1::varchar[])
`

func (q *Queries) GetOrdersByStatus(ctx context.Context, status []string) ([]Order, error) {
	rows, err := q.db.Query(ctx, getOrdersByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.PoolID,
			&i.ParentID,
			&i.Wallet,
			&i.Status,
			&i.Side,
			&i.Type,
			&i.Price,
			&i.Amount,
			&i.TwapAmount,
			&i.TwapParts,
			&i.PartialFilledAt,
			&i.FilledAt,
			&i.CancelledAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrdersByWallet = `-- name: GetOrdersByWallet :many
SELECT o1.id, o1.pool_id, o1.parent_id, o1.wallet, o1.status, o1.side, o1.type, o1.price, o1.amount, o1.twap_amount, o1.twap_parts, o1.partial_filled_at, o1.filled_at, o1.cancelled_at, o1.created_at, o2.id, o2.pool_id, o2.parent_id, o2.wallet, o2.status, o2.side, o2.type, o2.price, o2.amount, o2.twap_amount, o2.twap_parts, o2.partial_filled_at, o2.filled_at, o2.cancelled_at, o2.created_at FROM orders AS o1
LEFT JOIN orders AS o2 ON o1.id = o2.parent_id AND o2.parent_id IS NOT NULL
WHERE o1.wallet = $1
ORDER BY o1.created_at DESC
LIMIT $2 OFFSET $3
`

type GetOrdersByWalletParams struct {
	Wallet pgtype.Text `json:"wallet"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

type GetOrdersByWalletRow struct {
	ID              int64              `json:"id"`
	PoolID          string             `json:"pool_id"`
	ParentID        pgtype.Int8        `json:"parent_id"`
	Wallet          pgtype.Text        `json:"wallet"`
	Status          OrderStatus        `json:"status"`
	Side            OrderSide          `json:"side"`
	Type            OrderType          `json:"type"`
	Price           pgtype.Numeric     `json:"price"`
	Amount          pgtype.Numeric     `json:"amount"`
	TwapAmount      pgtype.Numeric     `json:"twap_amount"`
	TwapParts       pgtype.Int4        `json:"twap_parts"`
	PartialFilledAt pgtype.Timestamptz `json:"partial_filled_at"`
	FilledAt        pgtype.Timestamptz `json:"filled_at"`
	CancelledAt     pgtype.Timestamptz `json:"cancelled_at"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	Order           Order              `json:"order"`
}

func (q *Queries) GetOrdersByWallet(ctx context.Context, arg GetOrdersByWalletParams) ([]GetOrdersByWalletRow, error) {
	rows, err := q.db.Query(ctx, getOrdersByWallet, arg.Wallet, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrdersByWalletRow{}
	for rows.Next() {
		var i GetOrdersByWalletRow
		if err := rows.Scan(
			&i.ID,
			&i.PoolID,
			&i.ParentID,
			&i.Wallet,
			&i.Status,
			&i.Side,
			&i.Type,
			&i.Price,
			&i.Amount,
			&i.TwapAmount,
			&i.TwapParts,
			&i.PartialFilledAt,
			&i.FilledAt,
			&i.CancelledAt,
			&i.CreatedAt,
			&i.Order.ID,
			&i.Order.PoolID,
			&i.Order.ParentID,
			&i.Order.Wallet,
			&i.Order.Status,
			&i.Order.Side,
			&i.Order.Type,
			&i.Order.Price,
			&i.Order.Amount,
			&i.Order.TwapAmount,
			&i.Order.TwapParts,
			&i.Order.PartialFilledAt,
			&i.Order.FilledAt,
			&i.Order.CancelledAt,
			&i.Order.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertOrder = `-- name: InsertOrder :one
INSERT INTO orders (parent_id, wallet, pool_id, side, status, type, price, amount, twap_amount, twap_parts, filled_at, partial_filled_at, cancelled_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING id, pool_id, parent_id, wallet, status, side, type, price, amount, twap_amount, twap_parts, partial_filled_at, filled_at, cancelled_at, created_at
`

type InsertOrderParams struct {
	ParentID        pgtype.Int8        `json:"parent_id"`
	Wallet          pgtype.Text        `json:"wallet"`
	PoolID          string             `json:"pool_id"`
	Side            OrderSide          `json:"side"`
	Status          OrderStatus        `json:"status"`
	Type            OrderType          `json:"type"`
	Price           pgtype.Numeric     `json:"price"`
	Amount          pgtype.Numeric     `json:"amount"`
	TwapAmount      pgtype.Numeric     `json:"twap_amount"`
	TwapParts       pgtype.Int4        `json:"twap_parts"`
	FilledAt        pgtype.Timestamptz `json:"filled_at"`
	PartialFilledAt pgtype.Timestamptz `json:"partial_filled_at"`
	CancelledAt     pgtype.Timestamptz `json:"cancelled_at"`
}

func (q *Queries) InsertOrder(ctx context.Context, arg InsertOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, insertOrder,
		arg.ParentID,
		arg.Wallet,
		arg.PoolID,
		arg.Side,
		arg.Status,
		arg.Type,
		arg.Price,
		arg.Amount,
		arg.TwapAmount,
		arg.TwapParts,
		arg.FilledAt,
		arg.PartialFilledAt,
		arg.CancelledAt,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.PoolID,
		&i.ParentID,
		&i.Wallet,
		&i.Status,
		&i.Side,
		&i.Type,
		&i.Price,
		&i.Amount,
		&i.TwapAmount,
		&i.TwapParts,
		&i.PartialFilledAt,
		&i.FilledAt,
		&i.CancelledAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateOrder = `-- name: UpdateOrder :one
UPDATE orders
SET
    status = COALESCE($2, status),
    twap_amount = COALESCE($3, twap_amount),
    filled_at = COALESCE($4, filled_at),
    cancelled_at = COALESCE($5, cancelled_at),
    twap_amount = COALESCE($6, twap_amount),
    partial_filled_at = COALESCE($7, partial_filled_at)
WHERE id = $1
RETURNING id, pool_id, parent_id, wallet, status, side, type, price, amount, twap_amount, twap_parts, partial_filled_at, filled_at, cancelled_at, created_at
`

type UpdateOrderParams struct {
	ID              int64              `json:"id"`
	Status          OrderStatus        `json:"status"`
	TwapAmount      pgtype.Numeric     `json:"twap_amount"`
	FilledAt        pgtype.Timestamptz `json:"filled_at"`
	CancelledAt     pgtype.Timestamptz `json:"cancelled_at"`
	TwapAmount_2    pgtype.Numeric     `json:"twap_amount_2"`
	PartialFilledAt pgtype.Timestamptz `json:"partial_filled_at"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, updateOrder,
		arg.ID,
		arg.Status,
		arg.TwapAmount,
		arg.FilledAt,
		arg.CancelledAt,
		arg.TwapAmount_2,
		arg.PartialFilledAt,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.PoolID,
		&i.ParentID,
		&i.Wallet,
		&i.Status,
		&i.Side,
		&i.Type,
		&i.Price,
		&i.Amount,
		&i.TwapAmount,
		&i.TwapParts,
		&i.PartialFilledAt,
		&i.FilledAt,
		&i.CancelledAt,
		&i.CreatedAt,
	)
	return i, err
}
