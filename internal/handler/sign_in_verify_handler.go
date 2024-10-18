package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

// SignInVerifyHandler godoc
// @Summary      Окончание регистрации нового пользователя
// @Description  Окончание регистрации нового пользователя. Пользователь должен прислать токен для подтверждения его регистрации
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body model.SignInVerifyV1Request true "SignInVerifyData Request"
// @Success      200   {object}  model.TokensResponse "Registration successfully"
// @Failure      400   {object}  model.ErrorResponse "Invalid request body"
// @Failure      500   {object}  model.ErrorResponse "Internal server error"
// @Router       /v2/sign-in-verify [post]
func (h *Handler) SignInVerifyHandler(w http.ResponseWriter, r *http.Request) {
	var signInVerifyV1Request model.SignInVerifyV1Request

	if err := json.NewDecoder(r.Body).Decode(&signInVerifyV1Request); err != nil {
		log.Error().Err(err).Msgf("failed to decode sign in verify data, error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokens, err := h.auc.SignInVerifyV1(r.Context(), signInVerifyV1Request)
	if err != nil {
		log.Error().Err(err).Msgf("failed to sign in verify, error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(tokens); err != nil {
		log.Error().Err(err).Msgf("failed to encode tokens, error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
