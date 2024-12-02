package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// CreateAccountV1Handler godoc
// @Summary Добавление банковского счета
// @Description Добавление банковского счета
// @Tags Payment
// @Accept       json
// @Produce      json
// @Param        body  body model.CreateAccountV1Request true "CreateAccountV1Request"
// @Success      201   {object}  model.CreateAccountV1Response "Банковский счет успешно добавлен"
// @Failure      400   {object}  model.ErrorResponse "Invalid request body"
// @Failure      500   {object}  model.ErrorResponse "Internal server error"
// @Router       /v1/accounts [post].
func (h *Handler) CreateAccountV1Handler(w http.ResponseWriter, r *http.Request) {
	var createAccountRequest model.CreateAccountV1Request
	if err := json.NewDecoder(r.Body).Decode(&createAccountRequest); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in createAccountRequest, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if err := createAccountRequest.Validate(userID); err != nil {
		log.Error().Msgf("failed to validate createAccountRequest: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	response, err := h.puc.CreateAccountV1UseCase(r.Context(), createAccountRequest, userID)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.GetCode(), err.GetMessage(), err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if encodeError := json.NewEncoder(w).Encode(response); encodeError != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("SignInV1Handler EncodeError: %v", encodeError)
		h.HandleError(w, http.StatusInternalServerError, encodeError.Error())
		return
	}
}
