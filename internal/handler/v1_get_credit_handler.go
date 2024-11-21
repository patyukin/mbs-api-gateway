package handler

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) GetCreditV1(w http.ResponseWriter, r *http.Request) {
	creditID := r.PathValue("id")
	if creditID == "" {
		log.Error().Msg("r.PathValue(\"id\") missing creditID in GetCreditV1")
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("GetCreditV1 missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	response, err := h.cuc.GetCreditUseCase(r.Context(), creditID, userID)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.Code, err.Message, err.Description)
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
