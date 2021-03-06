package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		log.Printf("Accessing path %v with application %v from %v", r.RequestURI, userAgent, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}