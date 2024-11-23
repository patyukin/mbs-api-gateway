package handler

import (
	"encoding/json"
	"github.com/patyukin/mbs-api-gateway/internal/metrics"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) GetLogReportV1Handler(w http.ResponseWriter, r *http.Request) {
	metrics.TotalLogReport.Inc()
	var in model.GetLogReportV1Request

	userID := r.Header.Get(HeaderUserID)
	log.Info().Msgf("GetLogReportV1Handler userID: %v", userID)

	if userID == "" {
		metrics.FailedLogReport.Inc()
		log.Error().Msg("GetLogReportV1Handler missing userID")
		h.HandleError(w, http.StatusUnauthorized, "StatusUnauthorized")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("GetLogReportV1Handler DecodeError: %v", err)
		h.HandleError(w, http.StatusBadRequest, "invalid data")
		return
	}

	if err := in.Validate(); err != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("GetLogReportV1Handler ValidateError: %v", err.Description)
		h.HandleError(w, int(err.Code), err.Message)
		return
	}

	result, err := h.luc.GetLogReportV1UseCase(r.Context(), in)
	if err != nil {
		metrics.FailedLogReport.Inc()
		log.Error().Msgf("GetLogReportV1Handler UseCaseError: %v", err.Description)
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
