package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// SignInConfirmationHandler godoc
// @Summary      Окончание регистрации нового пользователя
// @Description  Окончание регистрации нового пользователя. Пользователь должен прислать токен для подтверждения его регистрации
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body model.SignInConfirmationV1Request true "SignInConfirmationData Request"
// @Success      200   {object}  model.TokensResponse "Registration successfully"
// @Failure      400   {object}  model.ErrorResponse "Invalid request body"
// @Failure      500   {object}  model.ErrorResponse "Internal server error"
// @Router       /v1/sign-in/confirmation [post].
func (h *Handler) SignInConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	var signInConfirmationV1Request model.SignInConfirmationV1Request
	if err := json.NewDecoder(r.Body).Decode(&signInConfirmationV1Request); err != nil {
		log.Error().Err(err).Msgf("failed to decode sign in verify data, error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokens, err := h.auc.SignInConfirmationV1UseCase(r.Context(), signInConfirmationV1Request)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, error: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(tokens); encodeErr != nil {
		log.Error().Msgf("failed to encode tokens, error: %v", encodeErr)
		h.HandleError(w, http.StatusInternalServerError, encodeErr.Error())
		return
	}
}
