package router

import (
	"fmt"
	_ "github.com/patyukin/mbs-api-gateway/docs"
	"github.com/patyukin/mbs-api-gateway/pkg/rate_limiter"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"net/http/pprof"
)

type Handler interface {
	Auth(next http.Handler) http.Handler
	CORS(next http.Handler) http.Handler
	LoggingMiddleware(next http.Handler) http.Handler
	RecoverMiddleware(next http.Handler) http.Handler
	RequestIDMiddleware(next http.Handler) http.Handler
	RateLimitMiddleware(limiter *rate_limiter.TokenBucketLimiter, next http.Handler) http.Handler
	HandleError(w http.ResponseWriter, code int, message string)
	HealthCheck(w http.ResponseWriter, r *http.Request)
	SignUpV1(w http.ResponseWriter, r *http.Request)
	SignInV1(w http.ResponseWriter, r *http.Request)
}

// Init docs
// @title Auth API
// @version 1.0
// @description Auth API for microservices
// @host http://0.0.0.0:5001
// @BasePath /
func Init(h Handler, limiter *rate_limiter.TokenBucketLimiter, srvAddress string) http.Handler {
	mux := http.NewServeMux()
	log.Info().Msgf("server address: %s", srvAddress)

	// swagger route
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://0.0.0.0%s/swagger/doc.json", srvAddress)),
	))

	// healthcheck
	mux.Handle("GET /healthcheck", http.HandlerFunc(h.HealthCheck))

	// api gateway routes
	mux.Handle("POST /v1/sign-up", http.HandlerFunc(h.SignUpV1))
	mux.Handle("POST /v1/sign-in", http.HandlerFunc(h.SignInV1))

	// pprof routes
	mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.Handle("/handle/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("goroutine"))

	// required middlewares
	withMiddlewareMux := h.RateLimitMiddleware(limiter, mux)
	withMiddlewareMux = h.RecoverMiddleware(mux)
	withMiddlewareMux = h.RequestIDMiddleware(withMiddlewareMux)
	withMiddlewareMux = h.LoggingMiddleware(withMiddlewareMux)
	withMiddlewareMux = h.CORS(withMiddlewareMux)

	return mux
}
