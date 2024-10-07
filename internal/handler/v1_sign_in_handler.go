package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) SignInV1(w http.ResponseWriter, r *http.Request) {
	metrics.TotalRegistrations.Inc()
	var in model.SignInV1Request

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1 DecodeError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1 ValidateError: %v", err)
		h.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.auc.SignInV1(r.Context(), in)
	if err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1 UseCaseError: %v", err)
		h.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	metrics.SuccessfulLogin.Inc()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err = json.NewEncoder(w).Encode(tokens); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1 EncodeError: %v", err)
		h.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
