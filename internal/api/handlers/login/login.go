package login

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golovpeter/passbox_backend/internal/common/auth_tokens"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/jmoiron/sqlx"
)

func Login(ctx *fiber.Ctx) error {
	conn := ctx.Locals("dbConn").(*sqlx.DB)

	var in In

	if err := ctx.BodyParser(&in); err != nil {
		return make_response.MakeInfoResponse(ctx, fiber.StatusUnprocessableEntity, 1, err.Error())
	}

	if in.Email == "" || in.Password == "" {
		return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "Incorrect data input")
	}

	elementExist := false
	err := conn.Get(&elementExist, "select exists(select email from users where email = $1)", in.Email)

	if err != nil {
		return make_response.MakeInfoResponse(ctx, fiber.StatusInternalServerError, 1, err.Error())
	}

	if !elementExist {
		return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "The user is not registered!")
	}

	var userId int
	err = conn.Get(&userId, "select id from users where email = $1", in.Email)

	tokenExist := false
	err = conn.Get(&tokenExist, "select exists(select id from tokens where id = $1)", userId)

	//FIXME: the nesting level is too deep
	if tokenExist {
		var tokens struct {
			accessToken  string `db:"access_token"`
			refreshToken string `db:"refresh_token"`
		}

		err = conn.QueryRowx(
			"select access_token, refresh_token from tokens where id = $1", userId,
		).Scan(&tokens.accessToken, &tokens.refreshToken)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusInternalServerError, 1, err.Error())
		}

		err = auth_tokens.ValidateToken(tokens.refreshToken)

		if err != nil && errors.Is(err, jwt.ErrTokenExpired) {
			newAccessToken, genErr := auth_tokens.GenerateJWT(in.Email)
			if genErr != nil {
				return err
			}

			newRefreshToken, genErr := auth_tokens.GenerateRefreshJWT()
			if genErr != nil {
				return err
			}

			_, genErr = conn.Exec("update tokens set access_token = $1, refresh_token = $2 where id = $3", newAccessToken, newRefreshToken, userId)

			response := Out{
				AccessToken: newAccessToken,
			}

			return ctx.JSON(response)
		}

		if err != nil {
			return err
		}

		response := Out{
			AccessToken: tokens.accessToken,
		}

		return ctx.JSON(response)
	}

	newAccessToken, err := auth_tokens.GenerateJWT(in.Email)
	if err != nil {
		return err
	}

	newRefreshToken, err := auth_tokens.GenerateRefreshJWT()
	if err != nil {
		return err
	}

	_, err = conn.Exec("insert into tokens values ($1, $2, $3)", userId, newAccessToken, newRefreshToken)

	response := Out{
		AccessToken: newAccessToken,
	}

	return ctx.JSON(response)
}
