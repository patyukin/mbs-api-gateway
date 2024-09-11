package router

import (
	"fmt"
	_ "github.com/patyukin/mbs-api-gateway/docs"
	"github.com/patyukin/mbs-api-gateway/internal/handler"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"net/http/pprof"
)

// Init godoc
// @title Auth API
// @version 1.0
// @description Api Gateway for MBS
// @host http://0.0.0.0:5001
// @BasePath /
func Init(h *handler.Handler, srvAddress string) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://0.0.0.0%s/swagger/doc.json", srvAddress)),
	))

	mux.Handle("POST /v1/sign-up", h.CORS(http.HandlerFunc(h.SignUpV1)))
	mux.Handle("POST /v1/sign-in", h.CORS(http.HandlerFunc(h.SignInV1)))

	// pprof
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

	recoveredMux := h.RecoverMiddleware(mux)

	return recoveredMux
}
