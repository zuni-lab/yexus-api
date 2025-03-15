package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/yexus-api/config"
	"github.com/zuni-lab/yexus-api/pkg/db"
)

type YieldMetricsResponse struct {
	Data []*db.YieldMetricData `json:"data"`
}

var httpClient = &http.Client{}

func FetchAndUpdateYieldMetrics(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, config.Env.YieldMetricsSource, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Use cached client to make request
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch yield metrics: %w", err)
	}
	defer resp.Body.Close()

	var yieldMetrics YieldMetricsResponse
	if err := json.NewDecoder(resp.Body).Decode(&yieldMetrics); err != nil {
		return fmt.Errorf("failed to decode yield metrics: %w", err)
	}

	_, err = db.DB.CreateBatchYieldMetricsTx(ctx, yieldMetrics.Data)
	if err != nil {
		return fmt.Errorf("failed to create batch yield metrics: %w", err)
	}

	log.Info().
		Int("count", len(yieldMetrics.Data)).
		Msg("Successfully updated yield metrics")

	return nil
}
