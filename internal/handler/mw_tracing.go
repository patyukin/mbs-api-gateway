package handler

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
)

func (h *Handler) TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tracer := otel.Tracer("http-server")

		ctx, span := tracer.Start(r.Context(), r.URL.Path)
		defer span.End()

		if requestUUID, ok := r.Context().Value(HeaderRequestUUID).(string); ok {
			span.SetAttributes(attribute.String("requestUUID", requestUUID))
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
