package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/config"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/evm"
	"github.com/zuni-lab/dexon-service/pkg/openobserve"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

func main() {
	config.LoadEnv()

	appName := config.Env.AppName + "-indexer"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	// Initialize services
	db.Init(ctx, config.Env.PostgresUrl, config.Env.MigrationUrl)
	defer db.Close()

	// Get pools from database
	pools, err := db.DB.GetPools(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get pools")
	}

	poolAddresses := utils.Map(pools, func(pool db.Pool) common.Address {
		return common.HexToAddress(pool.ID)
	})

	idxManager := evm.NewIndexerManager()
	idxManager.AddHandler(NewSwapHandler())

	if err := idxManager.Connect(); err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Ethereum client")
	}
	defer idxManager.Close()

	// Start indexing
	startBlock := uint64(config.Env.StartBlock)
	if err := idxManager.IndexPools(ctx, poolAddresses, startBlock); err != nil {
		log.Fatal().Err(err).Msg("Failed to index pools")
	}
}
