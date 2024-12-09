package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// CreditApplicationConfirmationV1Handler godoc
// @Summary Подтверждение заявки на кредит
// @Description Подтверждение заявки на кредит
// @Tags Credit
// @Accept  json
// @Produce json
// @Param   body body model.CreditApplicationConfirmationV1Request true "CreditApplicationConfirmationV1Request"
// @Success 201  {object}  model.CreditApplicationConfirmationV1Response "Заявка на кредит подтверждена"
// @Failure 400  {object} model.ErrorResponse "Invalid request body"
// @Failure 500  {object} model.ErrorResponse "Internal server error"
// @Router /v1/credit-applications/confirmation [post].
func (h *Handler) CreditApplicationConfirmationV1Handler(w http.ResponseWriter, r *http.Request) {
	var req model.CreditApplicationConfirmationV1Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in createAccountRequest, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("CreateAccountV1Handler missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	response, err := h.cuc.CreditApplicationConfirmationV1UseCase(r.Context(), req, userID)
	if err != nil {
		log.Error().Msgf(
			"failed to sign in verify, code: %d, message: %s, error: %v", err.GetCode(), err.GetMessage(), err.GetDescription(),
		)
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
