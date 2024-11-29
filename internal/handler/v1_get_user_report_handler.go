package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
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
		log.Error().Msgf("GetUserReportV1Handler ValidateError: %v", err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	report, reportError := h.ruc.GetUserReportV1UseCase(r.Context(), in)
	if reportError != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("failed h.ruc.GetUserReportV1UseCase: %v", reportError.GetDescription())
		h.HandleError(w, int(reportError.GetCode()), reportError.GetMessage())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(report); err != nil {
		log.Error().Err(err).Msgf("failed to encode result, error: %v", err)
		h.HandleError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
