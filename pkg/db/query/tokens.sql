-- name: CreateToken :one
INSERT INTO tokens (id, name, symbol, decimals, is_stable)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

