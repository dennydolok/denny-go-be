package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add((24 * time.Hour)).Unix()
	claims["email"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
