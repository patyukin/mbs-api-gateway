package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) GetUserReportV1Handler(w http.ResponseWriter, r *http.Request) {
	metrics.TotalLogReport.Inc()
	var in model.GetUserReportV1Request

	in.UserID = r.Header.Get(HeaderUserID)
	if in.UserID == "" {
		metrics.FailedLogReport.Inc()
		log.Error().Msg("GetUserReportV1Handler missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	in.StartDate = r.URL.Query().Get("start_date")
	in.EndDate = r.URL.Query().Get("end_date")
	if err := in.Validate(); err != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("GetUserReportV1Handler ValidateError: %v", err.Description)
		h.HandleError(w, int(err.Code), err.Message)
		return
	}

	report, reportError := h.ruc.GetUserReportV1UseCase(r.Context(), in)
	if reportError != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("failed h.ruc.GetUserReportV1UseCase: %v", reportError.Description)
		h.HandleError(w, int(reportError.Code), reportError.Message)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(report); err != nil {
		log.Error().Err(err).Msgf("failed to encode result, error: %v", err)
		h.HandleError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
