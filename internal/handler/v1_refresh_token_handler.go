package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// RefreshTokenV1Handler godoc
// @Summary      Обновление токена
// @Description  Обновление токена
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body model.RefreshTokenV1Request true "RefreshTokenData Request"
// @Success      200   {object}  model.RefreshTokenV1Response "Registration successfully"
// @Failure      400   {object}  model.ErrorResponse "Invalid request body"
// @Failure      500   {object}  model.ErrorResponse "Internal server error"
// @Router       /v1/refresh-token [post].
func (h *Handler) RefreshTokenV1Handler(w http.ResponseWriter, r *http.Request) {
	var refreshTokenRequest model.RefreshTokenV1Request

	if err := json.NewDecoder(r.Body).Decode(&refreshTokenRequest); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	response, err := h.auc.RefreshTokenV1UseCase(r.Context(), refreshTokenRequest)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.GetCode(), err.GetMessage(), err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.WriteHeader(http.StatusOK)
	if errEnc := json.NewEncoder(w).Encode(response); errEnc != nil {
		log.Error().Err(errEnc).Msgf("failed to encode tokens, error: %v", err)
		h.HandleError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
