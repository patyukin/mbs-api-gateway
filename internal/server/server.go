package server

import (
	"context"
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
		Addr:           addr,
		Handler:        s.h,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Duration(cfg.HttpServer.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.HttpServer.WriteTimeout) * time.Second,
	}

	log.Info().Msgf("Run server on %s", addr)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
