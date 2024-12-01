package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

// AddUserRoleV1Handler docs
// @Summary "Добавление роли пользователю"
// @Description "Добавление роли пользователю"
// @Tags User
// @Accept json
// @Produce json
// @Param AddUserRoleV1Request body model.AddUserRoleV1Request true "AddUserRoleV1Request"
// @Success 200 {object} model.AddUserRoleV1Response "Успешно добавлено"
// @Failure 400 {object} model.ErrorResponse "Invalid request"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /v1/users-roles [post]
func (h *Handler) AddUserRoleV1Handler(w http.ResponseWriter, r *http.Request) {
	var in model.AddUserRoleV1Request

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		log.Error().Err(err).Msgf("failed to decode sign in verify data, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		log.Error().Msgf("AddUserRoleV1Request ValidateError: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	response, err := h.auc.AddUserRoleV1UseCase(r.Context(), in)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, error: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if encodeError := json.NewEncoder(w).Encode(response); encodeError != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler EncodeError: %v", encodeError)
		h.HandleError(w, http.StatusInternalServerError, encodeError.Error())
		return
	}
}
