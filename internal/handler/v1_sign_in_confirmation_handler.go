package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// SignInConfirmationHandler godoc
// @Summary      Второй этап входа в систему
// @Description  Второй этап входа в систему по коду
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body model.SignInConfirmationV1Request true "SignInConfirmationData Request"
// @Success      200   {object}  model.SignInConfirmationV1Response "Registration successfully"
// @Failure      400   {object}  model.ErrorResponse "Invalid request body"
// @Failure      500   {object}  model.ErrorResponse "Internal server error"
// @Router       /v1/sign-in/confirmation [post].
func (h *Handler) SignInConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	var in model.SignInConfirmationV1Request
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		log.Error().Err(err).Msgf("failed to decode sign in verify data, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		log.Error().Msgf("SignInV1Handler ValidateError: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	tokens, err := h.auc.SignInConfirmationV1UseCase(r.Context(), in)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, error: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(tokens); encodeErr != nil {
		log.Error().Msgf("failed to encode tokens, error: %v", encodeErr)
		h.HandleError(w, http.StatusInternalServerError, encodeErr.Error())
		return
	}
}
