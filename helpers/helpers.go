package helpers

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func UserAuthorization(r *http.Request, userID uint) error {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return errors.New("unauthorized: No token cookie")
		}
		return errors.New("bad request: Failed to read token cookie")
	}

	tokenStr := cookie.Value
	claims := &jwt.StandardClaims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("pbi-rakamin"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New("unauthorized: Invalid token signature")
		}
		return errors.New("bad request: Failed to parse token")
	}

	if !tkn.Valid {
		return errors.New("unauthorized: Invalid token")
	}

	if claims.Subject != string(rune(userID)) {
		return errors.New("unauthorized: User ID does not match session User ID")
	}

	return nil
}
