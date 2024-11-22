package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) SignInV1Handler(w http.ResponseWriter, r *http.Request) {
	metrics.TotalRegistrations.Inc()
	var in model.SignInV1Request

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler DecodeError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler ValidateError: %v", err)
		h.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.auc.SignInV1(r.Context(), in)
	if err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("failed h.auc.SignInV1Handler: %v", err)
		h.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if response.Error != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler UseCaseError: %v", response.Error.Description)
		h.HandleError(w, int(response.Error.Code), response.Error.Message)
		return
	}

	metrics.SuccessfulLogin.Inc()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err = json.NewEncoder(w).Encode(response); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler EncodeError: %v", err)
		h.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
