package handler

import (
	"github.com/patyukin/mbs-api-gateway/pkg/rate_limiter"
	"net/http"
)

func (h *Handler) RateLimitMiddleware(limiter *rate_limiter.TokenBucketLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			h.HandleError(w, http.StatusTooManyRequests, "Too Many Requests")
			return
		}

		next.ServeHTTP(w, r)
	})
}
