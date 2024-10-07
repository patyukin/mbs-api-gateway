package handler

import (
	"context"
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	HeaderAuthorization = "Authorization"
	HeaderUserID        = "X-User-ID"
	HeaderRequestUUID   = "X-Request-UUID"
)

//go:generate go run github.com/vektra/mockery/v2@v2.45.1 --name=AuthUseCase
type AuthUseCase interface {
	SignUpV1(ctx context.Context, in model.SignUpV1Request) error
	SignInV1(ctx context.Context, in model.SignInV1Request) (model.SignInV1Response, error)
	SignInVerifyV1(ctx context.Context, in model.SignInVerifyV1Request) (model.SignInVerifyV1Response, error)
	GetJWTToken() []byte
}

type Handler struct {
	auc AuthUseCase
}

func New(auc AuthUseCase) *Handler {
	return &Handler{auc: auc}
}

func (h *Handler) HandleError(w http.ResponseWriter, code int, message string) {
	log.Error().Msgf("Error: %s", message)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(model.ErrorResponse{Error: message}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	metrics.IncomingTraffic.Inc()
	w.WriteHeader(http.StatusOK)
}
