package refresh_tokens

import (
	"errors"
	"github.com/golovpeter/passbox_backend/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golovpeter/passbox_backend/internal/common/auth_tokens"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/golovpeter/passbox_backend/internal/database"
)

// TODO: подумать, надо ли тут передавать access_token

func RefreshTokens(db database.Database, config *config.Config) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var in refreshIn

		if err := ctx.BodyParser(&in); err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnprocessableEntity, 1, err.Error())
		}

		err := auth_tokens.ValidateToken(in.RefreshToken)

		if err != nil && errors.Is(err, jwt.ErrTokenExpired) {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnauthorized, 1, "token is expired")
		}

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnauthorized, 1, err.Error())
		}

		claims, err := auth_tokens.GetTokenClaims(in.RefreshToken)

		if err != nil {
			return err
		}

		tokenExist := false
		err = db.ExistRefreshToken(&tokenExist, claims["DeviceID"].(string), in.RefreshToken)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		if !tokenExist {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnauthorized, 1, "no such refresh token")
		}

		newAccessToken, err := auth_tokens.GenerateJWT(
			config,
			int(claims["UserID"].(float64)),
			claims["Email"].(string),
			claims["DeviceID"].(string),
			auth_tokens.TokenTTL,
		)

		if err != nil {
			return err
		}

		newRefreshToken, err := auth_tokens.GenerateJWT(
			config,
			int(claims["UserID"].(float64)),
			claims["Email"].(string),
			claims["DeviceID"].(string),
			auth_tokens.RefreshTokenTTL,
		)

		if err != nil {
			return err
		}

		_, err = db.UpdateTokens(newAccessToken, newRefreshToken, claims["DeviceID"].(string))

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		out := refreshOut{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		}

		return ctx.JSON(out)
	}
}
