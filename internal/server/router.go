package server

import (
	"fmt"
	_ "github.com/patyukin/mbs-api-gateway/docs"
	"github.com/patyukin/mbs-api-gateway/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	RateLimitMiddleware(next http.Handler, rps float64, burst int) http.Handler
	TracingMiddleware(next http.Handler) http.Handler
	HandleError(w http.ResponseWriter, code int, message string)
	HealthCheck(w http.ResponseWriter, r *http.Request)
	SignUpV1(w http.ResponseWriter, r *http.Request)
	SignInV1(w http.ResponseWriter, r *http.Request)
	SignInVerifyHandler(w http.ResponseWriter, r *http.Request)
	RefreshTokenV1Handler(w http.ResponseWriter, r *http.Request)
	GetLogReportV1(w http.ResponseWriter, r *http.Request)
}

// Init docs
// @title Auth API
// @version 1.0
// @description Auth API for microservices
// @host http://0.0.0.0:5001
// @BasePath /
func Init(h Handler, cfg *config.Config, srvAddress string) http.Handler {
	mux := http.NewServeMux()

	// metrics
	mux.Handle("/metrics", promhttp.Handler())

	// swagger route
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://0.0.0.0%s/swagger/doc.json", srvAddress)),
	))

	// auth service routes
	mux.Handle("POST /v1/sign-up", http.HandlerFunc(h.SignUpV1))
	mux.Handle("POST /v1/sign-in", http.HandlerFunc(h.SignInV1))
	mux.Handle("POST /v1/sign-in-verify", http.HandlerFunc(h.SignInVerifyHandler))
	mux.Handle("POST /v1/refresh-token", http.HandlerFunc(h.RefreshTokenV1Handler))
	mux.Handle("POST /v1/users-roles", h.Auth(http.HandlerFunc(h.RefreshTokenV1Handler)))

	// payments service routes
	// mux.Handle("GET /v1/payments", http.HandlerFunc(h.GetPaymentV1))
	// mux.Handle("GET /v1/payments/{id}", http.HandlerFunc(h.GetPaymentByIDV1))
	// mux.Handle("POST /v1/payments", http.HandlerFunc(h.CreatePaymentV1))
	// mux.Handle("PUT /v1/payments/{id}", http.HandlerFunc(h.UpdatePaymentV1))
	// mux.Handle("DELETE /v1/payments/{id}", http.HandlerFunc(h.DeletePaymentV1)) - canceled

	// logger service routes
	mux.Handle("POST /v1/log-reports", h.Auth(http.HandlerFunc(h.GetLogReportV1)))

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
	withMiddlewareMux := h.TracingMiddleware(mux)
	withMiddlewareMux = h.LoggingMiddleware(withMiddlewareMux)
	withMiddlewareMux = h.CORS(withMiddlewareMux)
	withMiddlewareMux = h.RateLimitMiddleware(withMiddlewareMux, cfg.HttpServer.RateLimit.Rps, cfg.HttpServer.RateLimit.Burst)
	withMiddlewareMux = h.RequestIDMiddleware(withMiddlewareMux)
	withMiddlewareMux = h.RecoverMiddleware(withMiddlewareMux)

	return withMiddlewareMux
}
