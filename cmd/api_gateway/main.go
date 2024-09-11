package main

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/config"
	"github.com/patyukin/mbs-api-gateway/internal/handler"
	"github.com/patyukin/mbs-api-gateway/internal/server"
	"github.com/patyukin/mbs-api-gateway/internal/server/router"
	"github.com/patyukin/mbs-api-gateway/internal/usecase"
	rateLimiter "github.com/patyukin/mbs-api-gateway/pkg/rate_limiter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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

	uc := usecase.New([]byte(cfg.JwtSecret))
	h := handler.New(uc)

	// set limiter
	lmtr := rateLimiter.New(ctx, cfg.HttpServer.RateLimit, time.Second)

	r := router.Init(h, lmtr, srvAddress)
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
