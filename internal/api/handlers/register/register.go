package register

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/internal/common/hash_passwords"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/jmoiron/sqlx"
)

func Register(conn *sqlx.DB) func(ctx *fiber.Ctx) error {
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

		if elementExist {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "User already registered!")
		}

		passwordHash, err := hash_passwords.HashPassword(in.Password)
		if err != nil {
			return err
		}

		_, err = conn.Exec("insert into users (email, password_hash) values ($1, $2)", in.Email, passwordHash)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusInternalServerError, 1, err.Error())
		}

		return make_response.MakeInfoResponse(ctx, fiber.StatusOK, 0, "Registration was successful!")
	}
}
