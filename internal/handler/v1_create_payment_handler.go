package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) CreatePaymentV1Handler(w http.ResponseWriter, r *http.Request) {
	var createPaymentRequest model.CreatePaymentV1Request

	if err := json.NewDecoder(r.Body).Decode(&createPaymentRequest); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in createPaymentRequest, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("createPaymentRequest missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	createPaymentRequest.UserID = userID

	log.Debug().Msgf("createPaymentRequest: %v", createPaymentRequest)

	_, err := h.puc.CreatePaymentUseCase(r.Context(), createPaymentRequest)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.Code, err.Message, err.Description)
		h.HandleError(w, int(err.Code), err.Message)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
