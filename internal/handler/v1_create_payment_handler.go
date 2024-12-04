package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// CreatePaymentV1Handler godoc
// @Summary Создание платежа
// @Description Создание платежа
// @Tags Payment
// @Accept       json
// @Produce      json
// @Param        body  body model.CreatePaymentV1Request true "CreatePaymentData Request"
// @Success      200   {object}  model.CreatePaymentV1Request "Registration successfully"
// @Failure      400   {object}  model.ErrorResponse "Invalid request body"
// @Failure      500   {object}  model.ErrorResponse "Internal server error"
// @Router       /v1/create-payment [post].
func (h *Handler) CreatePaymentV1Handler(w http.ResponseWriter, r *http.Request) {
	var createPaymentRequest model.CreatePaymentV1Request
	if err := json.NewDecoder(r.Body).Decode(&createPaymentRequest); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in createPaymentRequest, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if err := createPaymentRequest.Validate(userID); err != nil {
		log.Error().Msgf("failed to validate createAccountRequest: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	result, err := h.puc.CreatePaymentV1UseCase(r.Context(), &createPaymentRequest, userID)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.GetCode(), err.GetMessage(), err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if encodeError := json.NewEncoder(w).Encode(result); encodeError != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler EncodeError: %v", encodeError)
		h.HandleError(w, http.StatusInternalServerError, encodeError.Error())
		return
	}
}
