package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func readBody(r *http.Request) (string, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(r.Body, &buf)

	body, err := io.ReadAll(tee)
	if err != nil {
		return "", err
	}

	r.Body = io.NopCloser(&buf)

	redactedBody, err := redactPassword(body)
	if err != nil {
		return string(body), nil
	}

	return redactedBody, nil
}

func redactPassword(body []byte) (string, error) {
	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if _, ok := data["password"]; ok {
		data["password"] = "***"
	}

	redactedBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(redactedBody), nil
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

		requestID := "unknown-request-id"
		if tempRequestID, ok := r.Context().Value(RequestID).(string); ok {
			requestID = tempRequestID
		}

		traceID := "unknown-trace-id"
		if tempTraceID, ok := r.Context().Value(TraceID).(string); ok {
			traceID = tempTraceID
		}

		log.Info().Msgf("Request ID: %s, Trace ID: %s, Path: %s, Query Params: %s, Body: %s", requestID, traceID, path, queryParams, body)

		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)

		log.Info().Msgf("Completed Request ID: %s, Trace ID: %s in %v", requestID, traceID, duration)
	})
}
