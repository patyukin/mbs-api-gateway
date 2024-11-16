package handler

import (
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func (h *Handler) TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//ctx, span := otel.Tracer(tracer.ProviderNameApiGateway).Start(r.Context(), "request")
		//ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", span.SpanContext().TraceID().String())
		//
		//defer span.End()

		span := opentracing.GlobalTracer().StartSpan("start-request")
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(r.Context(), span)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
