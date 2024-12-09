package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// ConfirmationPaymentV1Handler godoc
// @Summary Подтверждение платежа
// @Description Подтверждение платежа
// @Tags Payment
// @Accept       json
// @Produce      json
// @Param        body  body model.ConfirmationPaymentV1Request true "ConfirmationPaymentData Request"
// @Success      200   {object}  model.ConfirmationPaymentV1Request "Registration successfully"
// @Failure      400   {object}  model.ErrorResponse "Invalid request body"
// @Failure      500   {object}  model.ErrorResponse "Internal server error"
// @Router       /v1/confirmation-payment [post].
func (h *Handler) ConfirmationPaymentV1Handler(w http.ResponseWriter, r *http.Request) {
	var confirmationPaymentV1Request model.ConfirmationPaymentV1Request
	if err := json.NewDecoder(r.Body).Decode(&confirmationPaymentV1Request); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in confirmationPaymentV1Request, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if err := confirmationPaymentV1Request.Validate(userID); err != nil {
		log.Error().Msgf("failed to validate confirmationPaymentV1Request: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	result, err := h.puc.ConfirmationPaymentV1UseCase(r.Context(), confirmationPaymentV1Request, userID)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.GetCode(), err.GetMessage(), err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if encodeError := json.NewEncoder(w).Encode(result); encodeError != nil {
		log.Error().Msgf("SignInV1Handler EncodeError: %v", encodeError)
		h.HandleError(w, http.StatusInternalServerError, encodeError.Error())
		return
	}
}
