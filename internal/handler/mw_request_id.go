package handler

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestUUID := uuid.New().String()
		ctx := context.WithValue(r.Context(), HeaderRequestUUID, requestUUID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
