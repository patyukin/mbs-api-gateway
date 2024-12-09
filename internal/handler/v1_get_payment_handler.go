package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetPaymentV1Handler(w http.ResponseWriter, r *http.Request) {
	var getPaymentV1Request model.GetPaymentV1Request
	if err := json.NewDecoder(r.Body).Decode(&getPaymentV1Request); err != nil {
		log.Error().Msgf("GetPaymentV1Handler DecodeError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid request")
		return
	}

	userID := r.Header.Get(HeaderUserID)
	if err := getPaymentV1Request.Validate(userID); err != nil {
		log.Error().Msgf("GetPaymentV1Handler ValidateError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid request")
		return
	}

	result, err := h.puc.GetPaymentV1UseCase(r.Context(), getPaymentV1Request, userID)
	if err != nil {
		log.Error().Msgf("GetPaymentV1Handler GetPaymentV1UseCaseError: %v", err)
		h.HandleError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusOK)
	if errEnc := json.NewEncoder(w).Encode(result); errEnc != nil {
		log.Error().Err(errEnc).Msgf("GetPaymentV1Handler EncodeError: %v", err)
		h.HandleError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
