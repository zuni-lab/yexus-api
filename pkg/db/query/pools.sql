-- name: GetPools :many
SELECT * FROM pools;

-- name: GetPool :one
SELECT * FROM pools
WHERE id = $1 LIMIT 1;

-- name: CreatePool :one
INSERT INTO pools (id, token0_id, token1_id)
VALUES ($1, $2, $3)
RETURNING *;

