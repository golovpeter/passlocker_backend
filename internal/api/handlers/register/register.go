package register

import (
	"1password_copy_project/internal/common/make_response"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Register(ctx *fiber.Ctx) error {
	conn := ctx.Locals("dbConn").(*sqlx.DB)

	var in RegisterIn

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

	if elementExist {
		return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "User already registered!")
	}

	_, err = conn.Exec("insert into users (email, password) values ($1, $2)", in.Email, in.Password)

	if err != nil {
		return make_response.MakeInfoResponse(ctx, fiber.StatusInternalServerError, 1, err.Error())
	}

	return nil
}
