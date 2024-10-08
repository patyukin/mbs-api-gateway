package handler

import (
	"golang.org/x/time/rate"
	"net/http"
)

func (h *Handler) RateLimitMiddleware(next http.Handler, rps float64, burst int) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(rps), burst)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := limiter.Wait(r.Context()); err != nil {
			h.HandleError(w, http.StatusTooManyRequests, "Too Many Requests")
			return
		}

		next.ServeHTTP(w, r)
	})
}
