package handler

import "net/http"

func (h *Handler) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "POST")
			w.Header().Add("Access-Control-Allow-Methods", "GET")
			w.Header().Add("Access-Control-Allow-Headers", "Authorization")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

			return
		}

		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
