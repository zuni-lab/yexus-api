-- name: CreateToken :one
INSERT INTO tokens (id, name, symbol, decimals)
VALUES ($1, $2, $3, $4)
RETURNING *;

