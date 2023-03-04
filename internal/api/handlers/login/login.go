package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/internal/common/auth_tokens"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func Login(conn *sqlx.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
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

		newDeviceID := uuid.New()
		newAccessToken, err := auth_tokens.GenerateJWT(in.Email, newDeviceID.String())
		if err != nil {
			return err
		}

		newRefreshToken, err := auth_tokens.GenerateRefreshJWT()
		if err != nil {
			return err
		}

		_, err = conn.Exec("insert into tokens values ($1, $2, $3, $4)", userId, newDeviceID, newAccessToken, newRefreshToken)

		response := Out{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		}

		return ctx.JSON(response)
	}
}
