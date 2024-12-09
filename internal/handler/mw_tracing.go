package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func (h *Handler) TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			span := opentracing.GlobalTracer().StartSpan("start-request")
			defer span.Finish()

			jaegerCtx, ok := span.Context().(jaeger.SpanContext)
			if !ok {
				h.HandleError(w, http.StatusInternalServerError, "Internal Server Error: failed to retrieve trace context")
				return
			}

			traceID := jaegerCtx.TraceID().String()
			ctx := context.WithValue(r.Context(), TraceID, traceID)

			requestID := uuid.New().String()
			ctx = context.WithValue(ctx, RequestID, requestID)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
