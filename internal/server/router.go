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
	SignUpV1Handler(w http.ResponseWriter, r *http.Request)
	SignInV1Handler(w http.ResponseWriter, r *http.Request)
	SignInConfirmationHandler(w http.ResponseWriter, r *http.Request)
	RefreshTokenV1Handler(w http.ResponseWriter, r *http.Request)
	GetUserByIDV1Handler(w http.ResponseWriter, r *http.Request)
	GetUsersV1Handler(w http.ResponseWriter, r *http.Request)
	GetLogReportV1Handler(w http.ResponseWriter, r *http.Request)
	CreatePaymentV1Handler(w http.ResponseWriter, r *http.Request)
	CreateAccountV1Handler(w http.ResponseWriter, r *http.Request)
	ConfirmationPaymentV1Handler(w http.ResponseWriter, r *http.Request)
	CreateCreditApplicationV1Handler(w http.ResponseWriter, r *http.Request)
	CreditApplicationConfirmationV1Handler(w http.ResponseWriter, r *http.Request)
	GetCreditApplicationV1Handler(w http.ResponseWriter, r *http.Request)
	UpdateCreditApplicationStatusV1Handler(w http.ResponseWriter, r *http.Request)
	GetCreditV1Handler(w http.ResponseWriter, r *http.Request)
	GetListUserCreditsV1Handler(w http.ResponseWriter, r *http.Request)
	GetPaymentScheduleV1Handler(w http.ResponseWriter, r *http.Request)
	GetUserReportV1Handler(w http.ResponseWriter, r *http.Request)
	GetPaymentV1Handler(w http.ResponseWriter, r *http.Request)
	UpdatePaymentStatusV1Handler(w http.ResponseWriter, r *http.Request)
	GetTransactionsByPaymentV1Handler(w http.ResponseWriter, r *http.Request)
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
	mux.Handle(
		"/swagger/", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://0.0.0.0%s/swagger/doc.json", srvAddress)),
		),
	)

	// auth service routes
	mux.Handle("POST /v1/sign-up", http.HandlerFunc(h.SignUpV1Handler))
	mux.Handle("POST /v1/sign-in", http.HandlerFunc(h.SignInV1Handler))
	mux.Handle("POST /v1/sign-in-verify", http.HandlerFunc(h.SignInConfirmationHandler))
	mux.Handle("POST /v1/refresh-token", http.HandlerFunc(h.RefreshTokenV1Handler))
	mux.Handle("POST /v1/users-roles", h.Auth(http.HandlerFunc(h.RefreshTokenV1Handler)))
	mux.Handle("GET /v1/users/{id}", h.Auth(http.HandlerFunc(h.GetUserByIDV1Handler)))
	mux.Handle("GET /v1/users", h.Auth(http.HandlerFunc(h.GetUsersV1Handler)))

	// payments service routes
	mux.Handle("POST /v1/accounts", h.Auth(http.HandlerFunc(h.CreateAccountV1Handler)))
	mux.Handle("POST /v1/payments", h.Auth(http.HandlerFunc(h.CreatePaymentV1Handler)))
	mux.Handle("POST /v1/payments/verify", h.Auth(http.HandlerFunc(h.ConfirmationPaymentV1Handler)))
	mux.Handle("GET /v1/payments/{id}", h.Auth(http.HandlerFunc(h.GetPaymentV1Handler)))
	mux.Handle("PATCH /v1/payments/{id}", h.Auth(http.HandlerFunc(h.UpdatePaymentStatusV1Handler)))
	mux.Handle("PATCH /v1/payments/{id}/transactions", h.Auth(http.HandlerFunc(h.GetTransactionsByPaymentV1Handler)))

	// credit service routes
	mux.Handle("POST /v1/credit-applications", h.Auth(http.HandlerFunc(h.CreateCreditApplicationV1Handler)))
	mux.Handle("POST /v1/credit-applications/confirmation", h.Auth(http.HandlerFunc(h.CreditApplicationConfirmationV1Handler)))
	mux.Handle("GET /v1/credit-applications/{id}", h.Auth(http.HandlerFunc(h.GetCreditApplicationV1Handler)))
	mux.Handle("PATCH /v1/credit-applications/{id}", h.Auth(http.HandlerFunc(h.UpdateCreditApplicationStatusV1Handler)))
	mux.Handle("GET /v1/credits/{id}", h.Auth(http.HandlerFunc(h.GetCreditV1Handler)))
	mux.Handle("GET /v1/credits", h.Auth(http.HandlerFunc(h.GetListUserCreditsV1Handler)))
	mux.Handle("GET /v1/credits/{id}/payment-schedule", h.Auth(http.HandlerFunc(h.GetPaymentScheduleV1Handler)))

	// logger service routes
	mux.Handle("POST /v1/log-reports", h.Auth(http.HandlerFunc(h.GetLogReportV1Handler)))

	// report service routes
	mux.Handle("GET /v1/reports", h.Auth(http.HandlerFunc(h.GetUserReportV1Handler)))

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
