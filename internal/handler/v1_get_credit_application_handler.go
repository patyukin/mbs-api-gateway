package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *Handler) GetCreditApplicationV1Handler(w http.ResponseWriter, r *http.Request) {
	applicationID := r.PathValue("id")
	if applicationID == "" {
		log.Error().Msg("r.PathValue(\"id\") missing userID")
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("CreateAccountV1Handler missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	response, err := h.cuc.GetCreditApplicationV1UseCase(r.Context(), applicationID, userID)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.GetCode(), err.GetMessage(), err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.WriteHeader(http.StatusOK)
	if errEnc := json.NewEncoder(w).Encode(response); errEnc != nil {
		log.Error().Err(errEnc).Msgf("failed to encode tokens, error: %v", err)
		h.HandleError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
