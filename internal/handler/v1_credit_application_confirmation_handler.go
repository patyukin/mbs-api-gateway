package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) CreditApplicationConfirmationV1Handler(w http.ResponseWriter, r *http.Request) {
	var req model.CreditApplicationConfirmationV1Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in createAccountRequest, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("CreateAccountV1Handler missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	response, err := h.cuc.CreditApplicationConfirmationUseCase(r.Context(), req, userID)
	if err != nil {
		log.Error().Msgf(
			"failed to sign in verify, code: %d, message: %s, error: %v", err.Code, err.Message, err.Description,
		)
		h.HandleError(w, int(err.Code), err.Message)
		return
	}

	w.WriteHeader(http.StatusOK)
	if errEnc := json.NewEncoder(w).Encode(response); errEnc != nil {
		log.Error().Err(errEnc).Msgf("failed to encode tokens, error: %v", err)
		h.HandleError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
