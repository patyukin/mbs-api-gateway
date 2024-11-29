package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/patyukin/mbs-api-gateway/internal/config"
	"github.com/patyukin/mbs-api-gateway/internal/handler"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/server"
	"github.com/patyukin/mbs-api-gateway/internal/usecase/auth"
	"github.com/patyukin/mbs-api-gateway/internal/usecase/credit"
	"github.com/patyukin/mbs-api-gateway/internal/usecase/logger"
	"github.com/patyukin/mbs-api-gateway/internal/usecase/payment"
	"github.com/patyukin/mbs-api-gateway/internal/usecase/report"
	"github.com/patyukin/mbs-pkg/pkg/grpc_client"
	authpb "github.com/patyukin/mbs-pkg/pkg/proto/auth_v1"
	creditpb "github.com/patyukin/mbs-pkg/pkg/proto/credit_v1"
	loggerpb "github.com/patyukin/mbs-pkg/pkg/proto/logger_v1"
	paymentpb "github.com/patyukin/mbs-pkg/pkg/proto/payment_v1"
	reportpb "github.com/patyukin/mbs-pkg/pkg/proto/report_v1"
	"github.com/patyukin/mbs-pkg/pkg/tracing"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
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

	level, err := zerolog.ParseLevel(cfg.MinLogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid log level")
	}

	zerolog.SetGlobalLevel(level)

	srvAddress := fmt.Sprintf("0.0.0.0:%d", cfg.HttpServer.Port)

	// register metrics
	err = metrics.RegisterMetrics()
	if err != nil {
		log.Fatal().Msgf("failed to register metrics: %v", err)
	}

	_, closer, err := tracing.InitJaeger(cfg.TracerHost, ServiceName)
	if err != nil {
		log.Fatal().Msgf("failed to initialize tracer: %v", err)
	}

	defer closer()

	// auth service init
	authConn, err := grpc_client.NewGRPCClientServiceConn(cfg.GRPC.AuthService)
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

	// payment service init
	paymentConn, err := grpc_client.NewGRPCClientServiceConn(cfg.GRPC.PaymentService)
	if err != nil {
		log.Fatal().Msgf("failed to connect to payment service: %v", err)
	}

	defer func(paymentConn *grpc.ClientConn) {
		if err = paymentConn.Close(); err != nil {
			log.Error().Msgf("failed to close auth service connection: %v", err)
		}
	}(paymentConn)

	paymentClient := paymentpb.NewPaymentServiceClient(paymentConn)
	paymentUseCase := payment.New(paymentClient)

	// credit service init
	creditConn, err := grpc_client.NewGRPCClientServiceConn(cfg.GRPC.CreditService)
	if err != nil {
		log.Fatal().Msgf("failed to connect to CreditService: %v", err)
	}

	defer func(creditConn *grpc.ClientConn) {
		if err = creditConn.Close(); err != nil {
			log.Error().Msgf("failed to close creditConn service connection: %v", err)
		}
	}(creditConn)

	creditClient := creditpb.NewCreditsServiceV1Client(creditConn)
	creditUseCase := credit.New(creditClient)

	// logger service init
	loggerConn, err := grpc_client.NewGRPCClientServiceConn(cfg.GRPC.LoggerService)
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

	// report service init
	reportConn, err := grpc_client.NewGRPCClientServiceConn(cfg.GRPC.ReportService)
	if err != nil {
		log.Fatal().Msgf("failed to connect to report service: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		if err = conn.Close(); err != nil {
			log.Error().Msgf("failed to close report service connection: %v", err)
		}
	}(reportConn)

	reportClient := reportpb.NewReportServiceClient(reportConn)
	reportUseCase := report.New(reportClient)

	h := handler.New(authUseCase, loggerUseCase, paymentUseCase, creditUseCase, reportUseCase)
	r := server.Init(h, cfg, srvAddress)
	srv := server.New(r)

	errCh := make(chan error)

	go func() {
		log.Info().Msgf("starting http server on %d", cfg.HttpServer.Port)
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
