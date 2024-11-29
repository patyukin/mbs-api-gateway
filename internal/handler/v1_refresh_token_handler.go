package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

func (h *Handler) RefreshTokenV1Handler(w http.ResponseWriter, r *http.Request) {
	var refreshTokenRequest model.RefreshTokenV1Request

	if err := json.NewDecoder(r.Body).Decode(&refreshTokenRequest); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	response, err := h.auc.RefreshTokenV1UseCase(r.Context(), refreshTokenRequest)
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
