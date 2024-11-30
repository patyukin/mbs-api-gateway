package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// SignInV1Handler docs
// @Summary Авторизация пользователя
// @Description Авторизация пользователя в системе
// @Tags Auth
// @Accept json
// @Produce json
// @Param SignInRequest body model.SignInV1Request true "Данные для авторизации пользователя"
// @Success 200 "Пользователь успешно авторизован"
// @Failure 400 {object} model.ErrorResponse "Invalid request"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /v1/sign-in [post].
func (h *Handler) SignInV1Handler(w http.ResponseWriter, r *http.Request) {
	metrics.TotalLogin.Inc()
	var in model.SignInV1Request

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler DecodeError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler ValidateError: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	response, err := h.auc.SignInV1UseCase(r.Context(), in)
	if err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("failed h.auc.SignInV1Handler: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	metrics.SuccessfulLogin.Inc()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if encodeError := json.NewEncoder(w).Encode(response); encodeError != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler EncodeError: %v", encodeError)
		h.HandleError(w, http.StatusInternalServerError, encodeError.Error())
		return
	}
}
