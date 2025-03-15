-- name: GetYieldMetrics :many
SELECT * FROM yield_metrics
WHERE pool = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetLatestYieldMetric :one
SELECT * FROM yield_metrics
ORDER BY created_at DESC
LIMIT 1;

-- name: GetYieldMetricsForChat :many
WITH latest_metrics AS (
    SELECT DISTINCT ON (pool) *
    FROM yield_metrics
    WHERE project = ANY($1::text[])
    ORDER BY pool, created_at DESC
)
SELECT 
    pool,
    chain,
    project,
    symbol,
    apy,
    apy_base as "apyBase",
    apy_reward as "apyReward",
    reward_tokens as "rewardTokens",
    underlying_tokens as "underlyingTokens"
FROM latest_metrics
ORDER BY symbol ASC;