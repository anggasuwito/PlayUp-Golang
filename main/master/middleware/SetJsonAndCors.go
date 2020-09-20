package middleware

import (
	"net/http"
)

func SetJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("access-control-allow-origin", "*")
		next.ServeHTTP(w, r)
	})
}
