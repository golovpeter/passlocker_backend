package delete_password

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/internal/common/auth_tokens"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/golovpeter/passbox_backend/internal/common/parse_headers"
	"github.com/jmoiron/sqlx"
)

func DeletePassword(conn *sqlx.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var in deletePasswordIn

		accessToken, err := parse_headers.ParseAuthHeader(ctx)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnprocessableEntity, 1, err.Error())
		}

		if err = ctx.BodyParser(&in); err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnprocessableEntity, 1, err.Error())
		}

		claims, err := auth_tokens.GetTokenClaims(accessToken)

		if err != nil {
			return err
		}

		var passwordUserId int
		err = conn.Get(&passwordUserId, "select user_id from passwords where id = $1", in.PasswordID)

		if passwordUserId == 0 {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "there is no such note")
		}

		if passwordUserId != int(claims["UserID"].(float64)) {
			return make_response.MakeInfoResponse(
				ctx,
				fiber.StatusBadRequest,
				1,
				"this note belongs to another user",
			)
		}

		_, err = conn.Exec("delete from passwords where id = $1", in.PasswordID)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		return make_response.MakeInfoResponse(ctx, fiber.StatusOK, 0, "note successful deleted")
	}
}
