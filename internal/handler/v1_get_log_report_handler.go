package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) GetLogReportV1(w http.ResponseWriter, r *http.Request) {
	metrics.TotalLogReport.Inc()
	var in model.GetLogReportV1Request

	userID := r.Header.Get(HeaderUserID)
	log.Info().Msgf("GetLogReportV1 userID: %v", userID)

	if userID == "" {
		metrics.FailedLogReport.Inc()
		log.Error().Msg("GetLogReportV1 missing userID")
		h.HandleError(w, http.StatusUnauthorized, "StatusUnauthorized")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("GetLogReportV1 DecodeError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("GetLogReportV1 ValidateError: %v", err)
		h.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.luc.GetLogReport(r.Context(), in); err != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("GetLogReportV1 UseCaseError: %v", err.Description)
		h.HandleError(w, int(err.Code), err.Message)
		return
	}

	metrics.SuccessfulLogReport.Inc()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
