package handler

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h *Handler) Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msgf("Admin Middleware, user role: %v", r.Header.Get(HeaderUserRole))
		if r.Header.Get(HeaderUserRole) != SysAdmin {
			h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
