package handler

import (
	"context"
	"encoding/json"
	dtoRequest "github.com/patyukin/mbs-api-gateway/internal/dto/request"
	"github.com/patyukin/mbs-api-gateway/internal/dto/response"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	HeaderAuthorization = "Authorization"
	HeaderUserID        = "X-User-ID"
	HeaderRequestUUID   = "X-Request-UUID"
)

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=UseCase
type UseCase interface {
	SignUpV1(ctx context.Context, in dtoRequest.SignUpV1Request) error
	GetJWTToken() []byte
}

type Handler struct {
	uc UseCase
}

func New(uc UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) HandleError(w http.ResponseWriter, code int, message string) {
	log.Error().Msgf("Error: %s", message)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response.ErrorResponse{Error: message}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	metrics.IncomingTraffic.Inc()
	w.WriteHeader(http.StatusOK)
}
