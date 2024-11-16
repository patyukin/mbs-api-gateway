package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) GetLogReportV1(w http.ResponseWriter, r *http.Request) {
	metrics.TotalRegistrations.Inc()
	var in model.GetLogReportV1Request

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("GetLogReportV1 DecodeError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("GetLogReportV1 ValidateError: %v", err)
		h.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.luc.GetLogReport(r.Context(), in); err != nil {
		metrics.FailedLogin.Inc()
		log.Error().Msgf("GetLogReportV1 UseCaseError: %v", err.Description)
		h.HandleError(w, int(err.Code), err.Message)
		return
	}

	metrics.SuccessfulLogin.Inc()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
