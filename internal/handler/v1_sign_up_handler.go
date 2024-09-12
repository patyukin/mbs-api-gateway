package handler

import (
	"encoding/json"
	dtoRequest "github.com/patyukin/mbs-api-gateway/internal/dto/request"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/rs/zerolog/log"
	"net/http"
)

// SignUpV1 docs
// @Summary Регистрация нового пользователя
// @Description Регистрация нового пользователя в системе
// @Tags Auth
// @Accept json
// @Produce json
// @Param SignUpRequest body request.SignUpV1Request true "Запрос пользователя на регистрацию"
// @Success 201 "Пользователь успешно зарегистрирован"
// @Failure 400 {object} response.ErrorResponse "Invalid request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /v1/sign-up [post]
func (h *Handler) SignUpV1(w http.ResponseWriter, r *http.Request) {
	metrics.TotalRegistrations.Inc()
	var in dtoRequest.SignUpV1Request

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		metrics.FailedRegistrations.Inc()
		log.Error().Msgf("SignUpV1 DecodeError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		metrics.FailedRegistrations.Inc()
		log.Error().Msgf("SignUpV1 ValidateError: %v", err)
		h.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.uc.SignUpV1(r.Context(), in); err != nil {
		metrics.FailedRegistrations.Inc()
		log.Error().Msgf("SignUpV1 UseCaseError: %v", err)
		h.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	metrics.SuccessfulRegistrations.Inc()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
