package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// GetPaymentScheduleV1Handler godoc
// @Summary Получение расписания платежей
// @Description Получение расписания платежей
// @Tags Credit
// @Accept  json
// @Produce json
// @Success 200  {object}  model.GetPaymentScheduleV1Response "Расписание платежей получено"
// @Failure 400  {object} model.ErrorResponse "Invalid request body"
// @Failure 500  {object} model.ErrorResponse "Internal server error"
// @Router /v1/credits/{id}/payment-schedule [get].
func (h *Handler) GetPaymentScheduleV1Handler(w http.ResponseWriter, r *http.Request) {
	creditID := r.PathValue("id")
	if creditID == "" {
		log.Error().Msg("r.PathValue(\"id\") missing creditID")
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("CreateAccountV1Handler missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	response, err := h.cuc.GetPaymentScheduleV1UseCase(r.Context(), userID, creditID)
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
