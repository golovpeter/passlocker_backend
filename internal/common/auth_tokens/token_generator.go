package auth_tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golovpeter/passbox_backend/internal/config"
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

func GenerateJWT(config *config.Config, userID int, email string, deviceID string, expTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(expTime)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		userID,
		email,
		deviceID,
	})

	return token.SignedString([]byte(config.JwtKey))
}
