package refresh_tokens

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golovpeter/passbox_backend/internal/common/auth_tokens"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/jmoiron/sqlx"
)

func RefreshTokens(conn *sqlx.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		refreshToken := ctx.Cookies("refresh_token")

		err := auth_tokens.ValidateToken(refreshToken)

		if err != nil && errors.Is(err, jwt.ErrTokenExpired) {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnauthorized, 1, "token is expired")
		}

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnauthorized, 1, err.Error())
		}

		claims, err := auth_tokens.GetTokenClaims(refreshToken)

		if err != nil {
			return err
		}

		tokenExist := false
		err = conn.Get(
			&tokenExist,
			"select exists(select refresh_token, device_id from tokens where device_id = $1 and refresh_token=$2)",
			claims["DeviceID"],
			refreshToken,
		)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		if !tokenExist {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnauthorized, 1, "no such refresh token")
		}

		newAccessToken, err := auth_tokens.GenerateJWT(
			claims["Email"].(string),
			claims["DeviceID"].(string),
		)

		if err != nil {
			return err
		}

		newRefreshToken, err := auth_tokens.GenerateRefreshJWT(
			claims["Email"].(string),
			claims["DeviceID"].(string),
		)

		if err != nil {
			return err
		}

		_, err = conn.Exec(
			"update tokens set access_token=$1, refresh_token=$2 where device_id = $3",
			newAccessToken,
			newRefreshToken,
			claims["DeviceID"].(string),
		)

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
