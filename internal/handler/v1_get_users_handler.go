package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// GetUsersV1Handler godoc
// @Summary Получить список пользователей
// @Description Получить список пользователей (для админов)
// @Tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param GetUsersV1Request body model.GetUsersV1Request true "GetUsersV1Request"
// @Success 200 {object} model.GetUsersV1Response
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /v1/users [get].
func (h *Handler) GetUsersV1Handler(w http.ResponseWriter, r *http.Request) {
	var in model.GetUsersV1Request
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		log.Error().Err(err).Msgf("failed to decode sign in verify data, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		log.Error().Msgf("GetUsersV1Request ValidateError: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	if in.Page == 0 {
		in.Page = 1
	}

	if in.Limit == 0 {
		in.Limit = 10
	}

	response, err := h.auc.GetUsersV1UseCase(r.Context(), in)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, error: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
		log.Error().Msgf("failed GetUserByIDV1Handler to encode response, err: %v", encodeErr)
		h.HandleError(w, http.StatusInternalServerError, encodeErr.Error())
		return
	}
}
