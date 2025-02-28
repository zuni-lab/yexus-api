-- name: GetBlockProcessingState :one
SELECT * FROM block_processing_state
WHERE pool_address = $1 AND is_backfill = $2;

-- name: UpsertBlockProcessingState :exec
INSERT INTO block_processing_state (
    pool_address,
    last_processed_block,
    is_backfill
) VALUES (
    $1, $2, $3
)
ON CONFLICT (pool_address, is_backfill) DO UPDATE
SET 
    last_processed_block = EXCLUDED.last_processed_block,
    updated_at = NOW();