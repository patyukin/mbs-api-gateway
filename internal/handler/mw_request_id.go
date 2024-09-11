package handler

import (
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderRequestUUID, uuid.New().String())

		next.ServeHTTP(w, r)
	})
}
