package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

// CreateCreditV1Handler godoc
// @Summary      Создание кредитной заявки
// @Description  Создание кредитной заявки
// @Tags         Credit
// @Accept       json
// @Produce      json
// @Param        body  body model.CreateCreditV1Request true "CreateCreditData Request"
// @Success      201   {object}  model.CreateCreditV1Response "Registration successfully"
// @Failure      400   {object}  model.ErrorResponse "Invalid request body"
// @Failure      500   {object}  model.ErrorResponse "Internal server error"
// @Router       /v1/credits [post].
func (h *Handler) CreateCreditV1Handler(w http.ResponseWriter, r *http.Request) {
	var in model.CreateCreditV1Request
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		log.Error().Err(err).Msgf("failed to decode refresh token data in CreateCreditV1Request, error: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if err := in.Validate(userID); err != nil {
		log.Error().Msgf("failed to validate createAccountRequest: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	response, err := h.cuc.CreateCreditV1UseCase(r.Context(), in, userID)
	if err != nil {
		log.Error().Msgf(
			"failed CreateCreditApplicationV1UseCase, code: %d, message: %s, error: %v",
			err.GetCode(),
			err.GetMessage(),
			err.GetDescription(),
		)
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if encodeError := json.NewEncoder(w).Encode(response); encodeError != nil {
		log.Error().Msgf("CreateCreditApplicationV1Handler EncodeError: %v", encodeError)
		h.HandleError(w, http.StatusInternalServerError, encodeError.Error())
		return
	}

}
