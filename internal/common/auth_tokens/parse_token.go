package auth_tokens

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(inputToken string) error {
	_, err := jwt.Parse(inputToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func GetTokenClaims(inputToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(inputToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	return claims, nil
}
