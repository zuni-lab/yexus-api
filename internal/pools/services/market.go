package services

import (
	"context"

	"github.com/zuni-lab/dexon-service/pkg/db"
)

type GetMarketDataParams struct {
	PoolID     string `json:"pool_id" validate:"eth_addr"`
	TimeBucket string `json:"time_bucket" validate:"oneof=1s 1m 5m 15m 30m 1h 4h 6h 12h 24h"`
}

func GetMarketData(ctx context.Context, input GetMarketDataParams) ([]db.GetMarketDataRow, error) {
	return db.DB.GetMarketData(ctx, db.GetMarketDataParams{
		PoolID:     input.PoolID,
		TimeBucket: input.TimeBucket,
	})
}
