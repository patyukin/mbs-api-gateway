package handler

import (
	"bytes"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"time"
)

func readBody(r *http.Request) (string, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(r.Body, &buf)

	body, err := io.ReadAll(tee)
	if err != nil {
		return "", err
	}

	r.Body = io.NopCloser(&buf)

	return string(body), nil
}

func (h *Handler) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		queryParams := r.URL.RawQuery

		body, err := readBody(r)
		if err != nil {
			log.Info().Msgf("unable to read request body in LoggingMiddleware: %v", err)
			h.HandleError(w, http.StatusInternalServerError, "Unable to read request body")
			return
		}

		requestUUID := "request-uuid"
		if tempRequestUUID, ok := r.Context().Value(HeaderRequestUUID).(string); ok {
			requestUUID = tempRequestUUID
		}

		log.Info().Msgf("Request ID: %s, Path: %s, Query Params: %s, Body: %s", requestUUID, path, queryParams, body)

		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)

		log.Info().Msgf("Completed Request ID: %s in %v", requestUUID, duration)
	})
}
