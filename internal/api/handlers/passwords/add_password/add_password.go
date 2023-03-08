package add_password

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/internal/common/auth_tokens"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/golovpeter/passbox_backend/internal/common/parse_headers"
	"github.com/jmoiron/sqlx"
)

func AddPassword(conn *sqlx.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var in addPasswordIn

		accessToken, _ := parse_headers.ParseAuthHeader(ctx)

		if err := ctx.BodyParser(&in); err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnprocessableEntity, 1, err.Error())
		}

		//TODO: подумать над тем, какие поля будут обязательными, а какие нет
		if in.ServiceName == "" {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "Incorrect data input")
		}

		claims, err := auth_tokens.GetTokenClaims(accessToken)

		if err != nil {
			return err
		}

		_, err = conn.Exec(
			"insert into passwords (user_id, service_name, link, email, login, password) values ($1, $2, $3, $4, $5, $6)",
			claims["UserID"],
			in.Link,
			in.ServiceName,
			in.Email,
			in.Login,
			in.Password,
		)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		var passwordID int
		err = conn.Get(&passwordID, "select max(id) from passwords")

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		return ctx.JSON(fiber.Map{
			"error_code":  0,
			"password_id": passwordID,
		})
	}
}
