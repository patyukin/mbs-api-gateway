package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) CreateAccountV1(w http.ResponseWriter, r *http.Request) {
	var createAccountRequest model.CreateAccountV1Request

	if err := json.NewDecoder(r.Body).Decode(&createAccountRequest); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in createAccountRequest, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("CreateAccountV1 missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	createAccountRequest.UserID = userID
	_, err := h.puc.CreateAccountUseCase(r.Context(), createAccountRequest)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.Code, err.Message, err.Description)
		h.HandleError(w, int(err.Code), err.Message)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
