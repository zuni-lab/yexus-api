package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/cmd/worker/job"
	"github.com/zuni-lab/dexon-service/config"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/openobserve"
)

func main() {
	loadConfig()

	ctx := context.Background()

	db.Init(ctx, config.Env.PostgresUrl, config.Env.MigrationUrl)

	j := job.NewPriceTracker()

	err := j.Start(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start price tracker")
	}
}

func loadConfig() {
	config.LoadEnv()

	appName := config.Env.AppName
	if config.Env.IsDev {
		appName = appName + "-dev"
	}
	openobserve.Init(openobserve.OpenObserveConfig{
		Endpoint:    config.Env.OpenObserveEndpoint,
		Credential:  config.Env.OpenObserveCredential,
		ServiceName: appName,
		Env:         config.Env.Env,
	})

	config.InitLogger()
}
