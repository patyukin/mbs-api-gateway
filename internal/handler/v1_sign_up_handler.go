package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// SignUpV1Handler docs
// @Summary Регистрация нового пользователя
// @Description Регистрация нового пользователя в системе
// @Tags Auth
// @Accept json
// @Produce json
// @Param SignUpRequest body model.SignUpV1Request true "Данные для регистрации пользователя"
// @Success 201 "Пользователь успешно зарегистрирован"
// @Failure 400 {object} model.ErrorResponse "Invalid request"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /v1/sign-up [post].
func (h *Handler) SignUpV1Handler(w http.ResponseWriter, r *http.Request) {
	metrics.TotalRegistrations.Inc()
	var in model.SignUpV1Request

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		metrics.FailedRegistrations.Inc()
		log.Error().Msgf("SignUpV1Handler DecodeError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		metrics.FailedRegistrations.Inc()
		log.Error().Msgf("SignUpV1Handler ValidateError: %v", err)
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	msg, err := h.auc.SignUpV1UseCase(r.Context(), &in)
	if err != nil {
		metrics.FailedRegistrations.Inc()
		log.Error().Msgf("SignUpV1Handler UseCaseError: %v", err)
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	metrics.SuccessfulRegistrations.Inc()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if encodeErr := json.NewEncoder(w).Encode(msg); encodeErr != nil {
		metrics.FailedRegistrations.Inc()
		log.Error().Msgf("SignUpV1Handler EncodeError: %v", encodeErr)
		h.HandleError(w, http.StatusInternalServerError, encodeErr.Error())
		return
	}
}
