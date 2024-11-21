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

	result, err := h.luc.GetLogReport(r.Context(), in)
	if err != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("GetLogReportV1 UseCaseError: %v", err.Description)
		h.HandleError(w, int(err.Code), err.Message)
		return
	}

	log.Debug().Msgf("result.FileUrl: %v", result.FileUrl)

	w.WriteHeader(http.StatusOK)
	if errEnc := json.NewEncoder(w).Encode(result); errEnc != nil {
		log.Error().Err(errEnc).Msgf("failed to encode tokens, error: %v", err)
		h.HandleError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
