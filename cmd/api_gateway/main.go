package main

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/config"
	"github.com/patyukin/mbs-api-gateway/internal/handler"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/server"
	"github.com/patyukin/mbs-api-gateway/internal/usecase/auth"
	"github.com/patyukin/mbs-api-gateway/internal/usecase/logger"
	"github.com/patyukin/mbs-api-gateway/pkg/grpc_client"
	authpb "github.com/patyukin/mbs-pkg/pkg/proto/auth_v1"
	loggerpb "github.com/patyukin/mbs-pkg/pkg/proto/logger_v1"
	"github.com/patyukin/mbs-pkg/pkg/tracing"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const ServiceName = "ApiGateway"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("unable to load config: %v", err)
	}

	srvAddress := fmt.Sprintf(":%d", cfg.HttpServer.Port)

	// register metrics
	err = metrics.RegisterMetrics()
	if err != nil {
		log.Fatal().Msgf("failed to register metrics: %v", err)
	}

	_, closer, err := tracing.InitJaeger(fmt.Sprintf("localhost:6831"), ServiceName)
	if err != nil {
		log.Fatal().Msgf("failed to initialize tracer: %v", err)
	}

	defer closer()

	// auth service init
	authConn, err := grpc_client.NewGRPCClientServiceConn(cfg.GRPC.AuthServiceHost, cfg.GRPC.AuthServicePort)
	if err != nil {
		log.Fatal().Msgf("failed to connect to auth service: %v", err)
	}

	defer func(authConn *grpc.ClientConn) {
		if err = authConn.Close(); err != nil {
			log.Error().Msgf("failed to close auth service connection: %v", err)
		}
	}(authConn)

	authClient := authpb.NewAuthServiceClient(authConn)
	authUseCase := auth.New([]byte(cfg.JwtSecret), authClient)

	// logger service init
	loggerConn, err := grpc_client.NewGRPCClientServiceConn(cfg.GRPC.LoggerServiceHost, cfg.GRPC.LoggerServicePort)
	if err != nil {
		log.Fatal().Msgf("failed to connect to auth service: %v", err)
	}

	defer func(loggerConn *grpc.ClientConn) {
		if err = loggerConn.Close(); err != nil {
			log.Error().Msgf("failed to close auth service connection: %v", err)
		}
	}(loggerConn)

	loggerClient := loggerpb.NewLoggerServiceClient(loggerConn)
	loggerUseCase := logger.New(loggerClient)

	h := handler.New(authUseCase, loggerUseCase)
	r := server.Init(h, cfg, srvAddress)
	srv := server.New(r)

	errCh := make(chan error)

	go func() {
		log.Info().Msgf("starting server on %d", cfg.HttpServer.Port)
		if err = srv.Run(srvAddress, cfg); err != nil {
			log.Error().Msgf("failed starting server: %v", err)
			errCh <- err
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err = <-errCh:
		log.Error().Msgf("Failed to run, err: %v", err)
	case res := <-sigChan:
		if res == syscall.SIGINT || res == syscall.SIGTERM {
			log.Info().Msgf("Signal received")
		} else if res == syscall.SIGHUP {
			log.Info().Msgf("Signal received")
		}
	}

	log.Info().Msgf("Shutting Down")

	if err = srv.Shutdown(ctx); err != nil {
		log.Error().Msgf("failed server shutting down: %s", err.Error())
	}
}
