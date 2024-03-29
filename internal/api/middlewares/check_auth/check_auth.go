package check_auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golovpeter/passbox_backend/internal/common/auth_tokens"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/golovpeter/passbox_backend/internal/common/parse_headers"
	"github.com/golovpeter/passbox_backend/internal/database"
)

// TODO: подумать, надо ли тут передавать refresh_token

func CheckAuth(db database.Database) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {

		accessToken, err := parse_headers.ParseAuthHeader(ctx)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		jwtClaims, err := auth_tokens.GetTokenClaims(accessToken)

		if err != nil {
			return err
		}

		tokenExist := false
		err = db.ExistAccessToken(&tokenExist, jwtClaims["DeviceID"].(string), accessToken)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		if !tokenExist {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnauthorized, 1, "there are no such tokens")
		}

		err = auth_tokens.ValidateToken(accessToken)

		if err != nil && errors.Is(err, jwt.ErrTokenExpired) {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnauthorized, 1, "token is expired")
		}

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnauthorized, 1, err.Error())
		}

		return ctx.Next()
	}
}
