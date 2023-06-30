package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTAuth gets the Authorization header from the request and extracts the token. It then validates and either returns
// an error or continues with the request.
func JWTAuth(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// get bearer token
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeader[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		if validateToken(authHeaderParts[1]) {
			h(w, r)
		} else {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

	}
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("missionimpossible")
	
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}

		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
