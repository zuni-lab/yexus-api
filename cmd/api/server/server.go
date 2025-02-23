package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/config"
	"github.com/zuni-lab/dexon-service/pkg/db"
	"github.com/zuni-lab/dexon-service/pkg/openai"
	"github.com/zuni-lab/dexon-service/pkg/openobserve"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Server struct {
	Raw           *echo.Echo
	traceProvider *sdktrace.TracerProvider
}

func New() *Server {
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

	return &Server{e, tp}
}

func (s *Server) Start() error {
	ctx := context.Background()
	loadSvcs(ctx)
	s.printRoutes()

	return s.Raw.Start(fmt.Sprintf("%s:%s", config.Env.ApiHost, config.Env.Port))
}

func (s *Server) Close() {
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
}

func closeSvcs() {
	db.Close()
}
