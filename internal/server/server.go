package server

import (
	"context"
	"fmt"
	"github.com/patyukin/mbs-api-gateway/internal/config"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	h          http.Handler
}

func New(h http.Handler) *Server {
	return &Server{
		h: h,
	}
}

func (s *Server) Run(addr string, cfg *config.Config) error {
	s.httpServer = &http.Server{
		Addr:         addr,
		Handler:      s.h,
		ReadTimeout:  time.Duration(cfg.HTTPServer.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.HTTPServer.WriteTimeout) * time.Second,
	}

	log.Info().Msgf("Run server on %s", addr)

	return fmt.Errorf("failed to run server: %w", s.httpServer.ListenAndServe())
}

func (s *Server) Shutdown(ctx context.Context) error {
	return fmt.Errorf("failed to shutdown server: %w", s.httpServer.Shutdown(ctx))
}
