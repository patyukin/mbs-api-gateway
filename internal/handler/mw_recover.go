package handler

import (
	"github.com/patyukin/mbs-api-gateway/pkg/httperror"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().Msgf("Recovered from panic: %v", err)
				httperror.SendError(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
