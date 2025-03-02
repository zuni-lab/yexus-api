package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/config"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/evm"
	"github.com/zuni-lab/dexon-service/pkg/openai"
	"github.com/zuni-lab/dexon-service/pkg/openobserve"
	"github.com/zuni-lab/dexon-service/pkg/swap"
	"github.com/zuni-lab/dexon-service/pkg/utils"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Server struct {
	Raw           *echo.Echo
	traceProvider *sdktrace.TracerProvider
	rtManager     *evm.RealtimeManager
	ctx           context.Context
	cancel        context.CancelFunc
	wg            sync.WaitGroup
}

func New() *Server {
	ctx, cancel := context.WithCancel(context.Background())

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

	e := echo.New()
	e.HideBanner = true
	tp := openobserve.SetupTraceHTTP()

	setupAddHandlerEvent(e)
	setupMiddleware(e)
	setupErrorHandler(e)
	setupRoute(e)
	setupValidator(e)

	loadSvcs(ctx)

	return &Server{
		Raw:           e,
		traceProvider: tp,
		ctx:           ctx,
		cancel:        cancel,
	}
}

func (s *Server) Start() error {
	// Start realtime manager
	if err := s.startRealtimeManager(); err != nil {
		return err
	}

	s.printRoutes()

	log.Info().Msgf("🥪 Environment loaded: %+v", config.Env)

	return s.Raw.Start(fmt.Sprintf("%s:%s", config.Env.ApiHost, config.Env.Port))
}

func (s *Server) startRealtimeManager() error {
	pools, err := db.DB.GetPools(s.ctx)
	if err != nil {
		return fmt.Errorf("failed to get pools: %w", err)
	}

	poolAddresses := utils.Map(pools, func(pool db.Pool) common.Address {
		return common.HexToAddress(pool.ID)
	})

	s.rtManager = evm.NewRealtimeManager()
	s.rtManager.AddHandler(swap.NewSwapHandler())

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				if err := s.rtManager.Connect(); err != nil {
					log.Error().Err(err).Msg("Failed to connect to Ethereum client")
					continue
				}

				if err := s.rtManager.WatchPools(s.ctx, poolAddresses); err != nil {
					log.Error().Err(err).Msg("Error watching pools")
					s.rtManager.Close()
					continue
				}
			}
		}
	}()

	return nil
}

func (s *Server) Close() {
	s.cancel()  // Signal all goroutines to stop
	s.wg.Wait() // Wait for all goroutines to finish

	if s.rtManager != nil {
		s.rtManager.Close()
	}

	closeSvcs()
	s.Raw.Close()

	err := s.traceProvider.Shutdown(context.Background())
	if err != nil {
		log.Err(err).Msg("Error shutting down trace provider")
	}
}

func loadSvcs(ctx context.Context) {
	db.Init(ctx, config.Env.PostgresUrl, config.Env.MigrationUrl)
	openai.Init()
	swap.InitPoolInfo()
}

func closeSvcs() {
	db.Close()
}
