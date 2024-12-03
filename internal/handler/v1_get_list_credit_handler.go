package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
)

// GetListUserCreditsV1Handler godoc
// @Summary Получение списка кредитов
// @Description Получение списка кредитов
// @Tags Credit
// @Accept  json
// @Produce json
// @Success 200  {object}  model.GetListUserCreditsV1Response "Список кредитов получен"
// @Failure 400  {object} model.ErrorResponse "Invalid request body"
// @Failure 500  {object} model.ErrorResponse "Internal server error"
// @Router /v1/credits [get].
func (h *Handler) GetListUserCreditsV1Handler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get(HeaderUserID)
	if userID == "" {
		log.Error().Msg("GetCreditV1Handler missing userID")
		h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// parse limit, page
	var limit, page int32
	limitRequest := r.URL.Query().Get("limit")
	if limitRequest != "" {
		limit = int32(minLimit)
	} else {
		limitParsed, err := strconv.Atoi(limitRequest)
		if err != nil {
			log.Error().Msgf("failed to parse limit, error: %v", err)
			h.HandleError(w, http.StatusBadRequest, "invalid data")
			return
		}
		limit = int32(limitParsed)
	}

	pageRequest := r.URL.Query().Get("limit")
	if pageRequest != "" {
		page = int32(1)
	} else {
		pageParsed, err := strconv.ParseInt(pageRequest, 10, 32)
		if err != nil {
			log.Error().Msgf("failed to parse page, error: %v", err)
			h.HandleError(w, http.StatusBadRequest, "invalid data")
			return
		}
		page = int32(pageParsed)
	}

	response, err := h.cuc.GetListUserCreditsV1UseCase(
		r.Context(), model.GetListUserCreditsV1Request{
			UserID: userID,
			Page:   page,
			Limit:  limit,
		},
	)
	if err != nil {
		log.Error().Msgf("failed to sign in verify, code: %d, message: %s, error: %v", err.GetCode(), err.GetMessage(), err.GetDescription())
		h.HandleError(w, int(err.GetCode()), err.GetMessage())
		return
	}

	w.WriteHeader(http.StatusOK)
	if errEnc := json.NewEncoder(w).Encode(response); errEnc != nil {
		log.Error().Err(errEnc).Msgf("failed to encode tokens, error: %v", err)
		h.HandleError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
