package handler

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().Msgf("Recovered requestUUID: %v from panic: %v", r.Header.Get(HeaderRequestUUID), err)
				h.HandleError(w, http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
