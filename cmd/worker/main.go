package main

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/cmd/worker/handlers"
	"github.com/zuni-lab/dexon-service/config"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/evm"
	"github.com/zuni-lab/dexon-service/pkg/openobserve"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

func main() {
	loadConfig()

	ctx := context.Background()

	loadSvcs(ctx)

	mgr := evm.NewManager()

	mgr.AddHandler(handlers.NewSwapHandler())

	for {
		if err := mgr.Connect(); err != nil {
			log.Error().Err(err).Msg("Failed to connect to Ethereum client")
			continue
		}

		pools, err := db.DB.GetPools(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get pools")
			continue
		}

		poolAddresses := utils.Map(pools, func(pool db.Pool) common.Address {
			return common.HexToAddress(pool.ID)
		})

		if err := mgr.WatchPools(ctx, poolAddresses); err != nil {
			log.Error().Err(err).Msg("Error watching pools")
			mgr.Close()
			continue
		}
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

func loadSvcs(ctx context.Context) {
	db.Init(ctx, config.Env.PostgresUrl, config.Env.MigrationUrl)
}
