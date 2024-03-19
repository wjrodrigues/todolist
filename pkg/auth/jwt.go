package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"

	"github.com/golang-jwt/jwt/v5"
)

type JwtBody struct {
	ID string
}

func GenerateJWT(jwtBody JwtBody, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": jwtBody.ID,
	})

	value := make([]byte, 64)
	rand.Read(value)

	hmac := hmac.New(sha256.New, value)
	hmac.Write([]byte(secret))

	return token.SignedString(value)
}
