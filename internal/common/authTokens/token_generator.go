package authTokens

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenTTL        = time.Minute * 15
	refreshTokenTTL = time.Hour * 720
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Email string
}

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(tokenTTL)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		email,
	})

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func GenerateRefreshJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(refreshTokenTTL)},
	})

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
