package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) ConfirmationPaymentV1Handler(w http.ResponseWriter, r *http.Request) {
	var confirmationPaymentV1Request model.ConfirmationPaymentV1Request
	if err := json.NewDecoder(r.Body).Decode(&confirmationPaymentV1Request); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in confirmationPaymentV1Request, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("confirmationPaymentV1Request missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	confirmationPaymentV1Request.UserID = userID
	_, err := h.puc.ConfirmationPaymentV1UseCase(r.Context(), confirmationPaymentV1Request)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.Code, err.Message, err.Description)
		h.HandleError(w, int(err.Code), err.Message)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
