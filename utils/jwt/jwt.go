package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

type AuthConf struct {
	Secret string
}

func ValidateJWT(Secret string, r *http.Request) error {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return fmt.Errorf("missing token")
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.String(), "auth") {
			err := ValidateJWT(os.Getenv("Secret"), r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	}
}
