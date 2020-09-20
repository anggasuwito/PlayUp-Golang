package middleware

import (
	"log"
	"net/http"
	"playUp/main/master/utils/token"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getToken := r.Header.Get("token")
		if len(getToken) == 0 {
			http.Error(w, "Token Expired", http.StatusUnauthorized)
			return
		}
		validity, userName, _ := token.ValidateToken(getToken)
		log.Println(userName + " accessing " + r.RequestURI)
		if validity && userName == "admin" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid Sessions", http.StatusUnauthorized)
		}
	})
}
