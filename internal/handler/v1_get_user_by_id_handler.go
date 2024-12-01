package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

// GetUserByIDV1Handler godoc
// @Summary      Информация о пользователе
// @Description  Информация о пользователе
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        body  body model.GetUserByIDV1Request true "GetUserByIDV1Request"
// @Success      200   {object}  model.GetUserByIDV1Response "Registration successfully"
// @Failure      400   {object}  model.ErrorResponse "Invalid request body"
// @Failure      500   {object}  model.ErrorResponse "Internal server error"
// @Router       /v1/users/{id} [get].
func (h *Handler) GetUserByIDV1Handler(w http.ResponseWriter, r *http.Request) {
	var in model.GetUserByIDV1Request
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		log.Error().Err(err).Msgf("failed to decode sign in verify data, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userIDFromHeader := r.Header.Get(HeaderUserID)
	if userIDFromHeader == "" {
		log.Error().Msgf("GetUserByIDV1Handler missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	roleFromHeader := r.Header.Get(HeaderUserRole)
	if roleFromHeader == "" {
		log.Error().Msgf("GetUserByIDV1Handler missing role")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := in.Validate(userIDFromHeader, roleFromHeader); err != nil {
		log.Error().Msgf("GetUserByIDV1Request ValidateError: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	response, err := h.auc.GetUserByIDV1UseCase(r.Context(), in)
	if err != nil {
		log.Error().Msgf("failed GetUserByIDV1UseCase, err: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
		log.Error().Msgf("failed GetUserByIDV1Handler to encode response, err: %v", encodeErr)
		h.HandleError(w, http.StatusInternalServerError, encodeErr.Error())
		return
	}
}
