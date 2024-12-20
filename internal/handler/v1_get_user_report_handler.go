package handler

import (
	"encoding/json"
	"net/http"

	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// GetUserReportV1Handler godoc
// @Summary Получение отчета
// @Description Получение отчета
// @Tags Report
// @Accept  json
// @Produce json
// @Success 200  {object}  model.GetUserReportV1Response "Отчет получен"
// @Failure 400  {object} model.ErrorResponse "Invalid request body"
// @Failure 500  {object} model.ErrorResponse "Internal server error"
// @Router /v1/reports [get].
func (h *Handler) GetUserReportV1Handler(w http.ResponseWriter, r *http.Request) {
	metrics.TotalLogReport.Inc()
	var in model.GetUserReportV1Request

	in.UserID = r.Header.Get(HeaderUserID)
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
