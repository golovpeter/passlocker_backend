package auth_tokens

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenTTL        = time.Minute * 15
	RefreshTokenTTL = time.Hour * 720
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID   int
	Email    string
	DeviceID string
}

func GenerateJWT(userID int, email string, deviceID string, expTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(expTime)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		userID,
		email,
		deviceID,
	})

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
