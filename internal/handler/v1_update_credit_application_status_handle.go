package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

func (h *Handler) UpdateCreditApplicationStatusV1Handler(w http.ResponseWriter, r *http.Request) {
	applicationID := r.PathValue("id")
	if applicationID == "" {
		log.Error().Msg("r.PathValue(\"id\") missing applicationID in UpdateCreditApplicationV1")
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("CreateAccountV1Handler missing userID in UpdateCreditApplicationV1")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var updateCreditApplicationStatusV1Request model.UpdateCreditApplicationStatusV1Request
	if err := json.NewDecoder(r.Body).Decode(&updateCreditApplicationStatusV1Request); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in UpdateCreditApplicationStatusV1Request, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	updateCreditApplicationStatusV1Request.ApplicationID = applicationID
	if err := updateCreditApplicationStatusV1Request.Validate(); err != nil {
		log.Error().Msgf("failed to validate updateCreditApplicationStatusV1Request: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	response, err := h.cuc.UpdateCreditApplicationStatusV1UseCase(r.Context(), updateCreditApplicationStatusV1Request)
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