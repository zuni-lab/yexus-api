package db

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zuni-lab/yexus-api/pkg/utils"
)

type YieldMetricData struct {
	Pool             string          `json:"pool"`
	Chain            string          `json:"chain"`
	Project          string          `json:"project"`
	Symbol           string          `json:"symbol"`
	TvlUsd           float64         `json:"tvlUsd"`
	ApyBase          float64         `json:"apyBase"`
	ApyReward        float64         `json:"apyReward"`
	Apy              float64         `json:"apy"`
	RewardTokens     []string        `json:"rewardTokens"`
	ApyPct1d         float64         `json:"apyPct1D"`
	ApyPct7d         float64         `json:"apyPct7D"`
	ApyPct30d        float64         `json:"apyPct30D"`
	Stablecoin       bool            `json:"stablecoin"`
	IlRisk           string          `json:"ilRisk"`
	Exposure         string          `json:"exposure"`
	Predictions      json.RawMessage `json:"predictions"`
	PoolMeta         string          `json:"poolMeta"`
	UnderlyingTokens []string        `json:"underlyingTokens"`
	Il7d             float64         `json:"il7d"`
	ApyBase7d        float64         `json:"apyBase7d"`
	ApyMean30d       float64         `json:"apyMean30d"`
	VolumeUsd1d      float64         `json:"volumeUsd1d"`
	VolumeUsd7d      float64         `json:"volumeUsd7d"`
	ApyBaseInception float64         `json:"apyBaseInception"`
}

type CreateBatchYieldMetricsTxResult struct {
	RowsAffected int64
}

func (store *SqlStore) CreateBatchYieldMetricsTx(ctx context.Context, yieldMetrics []*YieldMetricData) (CreateBatchYieldMetricsTxResult, error) {
	var result CreateBatchYieldMetricsTxResult

	// Calculate max records per batch (65535 / 24 parameters per record)
	const maxParamsPerQuery = 65535
	const paramsPerRecord = 24
	const batchSize = maxParamsPerQuery / paramsPerRecord

	// Process in batches
	for i := 0; i < len(yieldMetrics); i += batchSize {
		end := i + batchSize
		if end > len(yieldMetrics) {
			end = len(yieldMetrics)
		}

		batchMetrics := yieldMetrics[i:end]

		conn, err := store.connPool.Acquire(ctx)
		if err != nil {
			return result, fmt.Errorf("acquire connection: %w", err)
		}

		// Build values portion of query for this batch
		valueStrings := make([]string, 0, len(batchMetrics))
		valueArgs := make([]interface{}, 0, len(batchMetrics)*paramsPerRecord)

		for j, metric := range batchMetrics {
			tvlUsd, err := scanNumeric(metric.TvlUsd, "tvl_usd")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan tvl_usd: %w", err)
			}

			apyBase, err := scanNumeric(metric.ApyBase, "apy_base")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan apy_base: %w", err)
			}

			apyReward, err := scanNumeric(metric.ApyReward, "apy_reward")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan apy_reward: %w", err)
			}

			apy, err := scanNumeric(metric.Apy, "apy")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan apy: %w", err)
			}

			apyPct1d, err := scanNumeric(metric.ApyPct1d, "apy_pct_1d")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan apy_pct_1d: %w", err)
			}

			apyPct7d, err := scanNumeric(metric.ApyPct7d, "apy_pct_7d")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan apy_pct_7d: %w", err)
			}

			apyPct30d, err := scanNumeric(metric.ApyPct30d, "apy_pct_30d")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan apy_pct_30d: %w", err)
			}

			stablecoin, err := scanBool(metric.Stablecoin, "stablecoin")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan stablecoin: %w", err)
			}

			ilRisk, err := scanString(metric.IlRisk, "il_risk")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan il_risk: %w", err)
			}

			exposure, err := scanString(metric.Exposure, "exposure")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan exposure: %w", err)
			}

			poolMeta, err := scanString(metric.PoolMeta, "pool_meta")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan pool_meta: %w", err)
			}

			il7d, err := scanNumeric(metric.Il7d, "il_7d")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan il_7d: %w", err)
			}

			apyBase7d, err := scanNumeric(metric.ApyBase7d, "apy_base_7d")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan apy_base_7d: %w", err)
			}

			apyMean30d, err := scanNumeric(metric.ApyMean30d, "apy_mean_30d")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan apy_mean_30d: %w", err)
			}

			volumeUsd1d, err := scanNumeric(metric.VolumeUsd1d, "volume_usd_1d")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan volume_usd_1d: %w", err)
			}

			volumeUsd7d, err := scanNumeric(metric.VolumeUsd7d, "volume_usd_7d")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan volume_usd_7d: %w", err)
			}

			apyBaseInception, err := scanNumeric(metric.ApyBaseInception, "apy_base_inception")
			if err != nil {
				conn.Release()
				return result, fmt.Errorf("failed to scan apy_base_inception: %w", err)
			}

			// Create placeholders for this row
			offset := j * paramsPerRecord
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
				offset+1, offset+2, offset+3, offset+4, offset+5, offset+6, offset+7, offset+8, offset+9, offset+10,
				offset+11, offset+12, offset+13, offset+14, offset+15, offset+16, offset+17, offset+18, offset+19, offset+20,
				offset+21, offset+22, offset+23, offset+24))

			valueArgs = append(valueArgs,
				metric.Pool, metric.Chain, metric.Project, metric.Symbol,
				tvlUsd, apyBase, apyReward, apy,
				metric.RewardTokens, apyPct1d, apyPct7d, apyPct30d,
				stablecoin, ilRisk, exposure, metric.Predictions,
				poolMeta, metric.UnderlyingTokens, il7d, apyBase7d,
				apyMean30d, volumeUsd1d, volumeUsd7d, apyBaseInception)
		}

		query := fmt.Sprintf(`
			INSERT INTO yield_metrics (
				pool, chain, project, symbol, tvl_usd, apy_base, apy_reward, apy,
				reward_tokens, apy_pct_1d, apy_pct_7d, apy_pct_30d, stablecoin,
				il_risk, exposure, predictions, pool_meta, underlying_tokens,
				il_7d, apy_base_7d, apy_mean_30d, volume_usd_1d, volume_usd_7d,
				apy_base_inception
			) VALUES %s`, strings.Join(valueStrings, ","))

		tag, err := conn.Exec(ctx, query, valueArgs...)
		conn.Release()

		if err != nil {
			return result, fmt.Errorf("executing bulk insert: %w", err)
		}

		result.RowsAffected += tag.RowsAffected()
	}

	return result, nil
}

func scanNumeric(value float64, field string) (*pgtype.Numeric, error) {
	val, err := utils.ScanNumericValue(fmt.Sprintf("%.18f", value))
	if err != nil {
		return nil, fmt.Errorf("failed to scan %s: %w", field, err)
	}
	return val, nil
}

func scanString(value string, field string) (*pgtype.Text, error) {
	val, err := utils.ScanStringValue(value)
	if err != nil {
		return nil, fmt.Errorf("failed to scan %s: %w", field, err)
	}
	return val, nil
}

func scanBool(value bool, field string) (*pgtype.Bool, error) {
	val, err := utils.ScanBoolValue(fmt.Sprintf("%t", value))
	if err != nil {
		return nil, fmt.Errorf("failed to scan %s: %w", field, err)
	}
	return val, nil
}
