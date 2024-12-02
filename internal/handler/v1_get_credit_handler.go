package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// GetCreditV1Handler godoc
// @Summary Получение информации о кредите
// @Description Получение информации о кредите
// @Tags Credit
// @Accept  json
// @Produce json
// @Success 200  {object}  model.GetCreditV1Response "Информация о кредите получена"
// @Failure 400  {object} model.ErrorResponse "Invalid request body"
// @Failure 500  {object} model.ErrorResponse "Internal server error"
// @Router /v1/credits/{id} [get].
func (h *Handler) GetCreditV1Handler(w http.ResponseWriter, r *http.Request) {
	creditID := r.PathValue("id")
	if creditID == "" {
		log.Error().Msg("r.PathValue(\"id\") missing creditID in GetCreditV1Handler")
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("GetCreditV1Handler missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	response, err := h.cuc.GetCreditV1UseCase(r.Context(), creditID, userID)
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
