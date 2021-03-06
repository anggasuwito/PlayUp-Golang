package token

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

func ValidateToken(tokenString string) (bool, string, error) {
	mySigningKey := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	var userName string
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		userName = claims.Audience
	} else {
		log.Println(`error => `, err)
	}
	return token.Valid, userName, err
}
