package auth_tokens

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
	Email    string
	DeviceID string
}

func GenerateJWT(email string, deviceID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(tokenTTL)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		email,
		deviceID,
	})

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func GenerateRefreshJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(refreshTokenTTL)},
	})

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
