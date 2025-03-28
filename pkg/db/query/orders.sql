-- name: InsertOrder :one
INSERT INTO orders (
    pool_ids, parent_id, wallet, status, side, type,
    price, amount, slippage, twap_interval_seconds,
    twap_executed_times, twap_current_executed_times,
    twap_min_price, twap_max_price, twap_started_at, deadline,
    signature, paths, nonce, tx_hash, actual_amount,
    partial_filled_at, filled_at, rejected_at,
    cancelled_at, created_at)
VALUES ($1, $2, $3, $4, $5, $6,
        $7, $8, $9, $10,
        $11, $12, $13,
        $14, $15, $16,
        $17, $18, $19, $20,
        $21, $22, $23, $24, $25, $26)
RETURNING
    id, pool_ids, parent_id, wallet, status, side, type,
    price, amount, slippage, twap_interval_seconds,
    twap_executed_times, twap_current_executed_times,
    twap_min_price, twap_max_price, twap_started_at, deadline, nonce,
    paths, tx_hash, actual_amount, partial_filled_at, filled_at, rejected_at,
    cancelled_at, created_at;

-- name: GetOrdersByWallet :many
SELECT id, pool_ids, parent_id, wallet, status, side, type,
       price, amount, actual_amount, slippage, twap_interval_seconds,
       twap_executed_times, twap_current_executed_times,
       twap_min_price, twap_max_price, twap_started_at, deadline, nonce,
       paths, tx_hash, actual_amount, partial_filled_at, filled_at, rejected_at,
       cancelled_at, created_at
FROM orders
WHERE wallet = $1
    AND (
        ARRAY_LENGTH(@status::order_status[], 1) IS NULL
        OR (
            status = ANY(@status)
            AND (
                status <> 'PENDING'
                OR deadline IS NULL
                OR deadline > NOW() --Skip expired orders
            )
        )
    )
    AND (
        ARRAY_LENGTH(@not_status::order_status[], 1) IS NULL
        OR (
            NOT status = ANY(@not_status)
            AND (
                status <> 'PENDING'
                OR deadline IS NULL
                OR deadline <= NOW()
            )
        )
    )
    AND (
        ARRAY_LENGTH(@types::order_type[], 1) IS NULL
        OR type = ANY(@types)
    )
    AND (
        sqlc.narg(side)::order_side IS NULL
        OR side = @side
    )
    AND (
        CASE
            WHEN sqlc.narg(parent_id)::bigint IS NULL THEN parent_id IS NULL
            ELSE parent_id = @parent_id
        END
    )
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountOrdersByWallet :one
SELECT COUNT(*) AS total_counts
FROM orders
WHERE wallet = $1
    AND (
        ARRAY_LENGTH(@status::order_status[], 1) IS NULL
        OR (
            status = ANY(@status)
            AND (
                status <> 'PENDING'
                OR deadline IS NULL
                OR deadline > NOW() --Skip expired orders
            )
        )
    )
    AND (
        ARRAY_LENGTH(@not_status::order_status[], 1) IS NULL
        OR (
            NOT status = ANY(@not_status)
            AND (
                status <> 'PENDING'
                OR deadline IS NULL
                OR deadline <= NOW()
            )
        )
    )
    AND (
        ARRAY_LENGTH(@types::order_type[], 1) IS NULL
        OR type = ANY(@types)
    )
    AND (
        sqlc.narg(side)::order_side IS NULL
        OR side = @side
    )
    AND (
        CASE
            WHEN sqlc.narg(parent_id)::bigint IS NULL THEN parent_id IS NULL
            ELSE parent_id = @parent_id
        END
    );

-- name: GetOrderByID :one
SELECT id, pool_ids, parent_id, wallet, status, side, type,
       price, amount, slippage, twap_interval_seconds,
       twap_executed_times, twap_current_executed_times,
       twap_min_price, twap_max_price, twap_started_at, deadline, nonce,
       paths, tx_hash, actual_amount, partial_filled_at, filled_at, rejected_at,
       cancelled_at, created_at
FROM orders
WHERE wallet = $1 AND id = $2;

-- name: GetMatchedOrder :one
SELECT * FROM orders
WHERE (
        (side = 'BUY' AND type = 'LIMIT' AND price >= $1)
        OR (side = 'SELL' AND type = 'LIMIT' AND price <= $1)
        OR (side = 'BUY' AND type = 'STOP' AND price <= $1)
        OR (side = 'SELL' AND type = 'STOP' AND price >= $1)
        OR (type = 'TWAP' AND price BETWEEN twap_min_price AND twap_max_price)
    )
    AND status IN ('PENDING', 'PARTIAL_FILLED')
    AND (
        type <> 'TWAP'
        OR ( -- Check TWAP condition
            twap_current_executed_times < twap_executed_times
            AND (
                twap_started_at IS NULL
                OR twap_started_at >= NOW()
            )
            AND (
                partial_filled_at IS NULL
                OR partial_filled_at + (twap_interval_seconds || ' seconds')::interval < NOW()
            )
        )
    )
    AND (
        deadline IS NULL
        OR deadline > NOW()
    )
ORDER BY created_at ASC
LIMIT 1;

-- name: GetMatchedTwapOrder :many
SELECT * FROM orders
WHERE type = 'TWAP'
  AND twap_min_price IS NULL
  AND (
    twap_started_at IS NULL
    OR twap_started_at <= NOW()
  )
  AND status IN ('PENDING', 'PARTIAL_FILLED')
  AND twap_current_executed_times < twap_executed_times
  AND (
        partial_filled_at IS NULL
        OR partial_filled_at + (twap_interval_seconds || ' seconds')::interval < NOW()
  );

-- name: CancelOrder :one
UPDATE orders
SET
    status = 'CANCELLED',
    cancelled_at = $1
WHERE id = $2 AND wallet = $3 AND status NOT IN ('REJECTED', 'FILLED')
RETURNING
    id, pool_ids, parent_id, wallet, status, side, type,
    price, amount, slippage, twap_interval_seconds,
    twap_executed_times, twap_current_executed_times,
    twap_min_price, twap_max_price, deadline, nonce,
    paths, tx_hash, actual_amount, partial_filled_at, filled_at, rejected_at,
    cancelled_at, created_at;

-- name: CancelAllOrders :exec
UPDATE orders
SET
    status = 'CANCELLED',
    cancelled_at = $1
WHERE wallet = $2 AND status NOT IN ('REJECTED', 'FILLED')
    RETURNING
    id, pool_ids, parent_id, wallet, status, side, type,
    price, amount, slippage, twap_interval_seconds,
    twap_executed_times, twap_current_executed_times,
    twap_min_price, twap_max_price, deadline, nonce,
    paths, tx_hash, actual_amount, partial_filled_at, filled_at, rejected_at,
    cancelled_at, created_at;

-- name: FillOrder :one
UPDATE orders
SET
    status = 'FILLED',
    filled_at = $1,
    tx_hash = $2,
    actual_amount = $3
WHERE id = $4
RETURNING *;

-- name: FillTwapOrder :one
UPDATE orders
SET
    status = $1,
    twap_current_executed_times = $2,
    partial_filled_at = COALESCE($3, partial_filled_at),
    filled_at = $4,
    actual_amount = COALESCE(actual_amount, 0) + $5
WHERE id = $6
RETURNING *;

-- name: RejectOrder :one
UPDATE orders
SET
    status = 'REJECTED',
    rejected_at = $1
WHERE id = $2
RETURNING *;
