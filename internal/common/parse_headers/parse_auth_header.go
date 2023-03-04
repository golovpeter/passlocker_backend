package parse_headers

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ParseAuthHeader(ctx *fiber.Ctx) (string, error) {
	headers := ctx.GetReqHeaders()

	authHeader := strings.Split(headers["Authorization"], " ")

	if len(authHeader) != 2 || authHeader[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if authHeader[1] == "" {
		return "", errors.New("token is empty")
	}

	return authHeader[1], nil
}
