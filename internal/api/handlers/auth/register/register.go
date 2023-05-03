package register

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/internal/common/hash_passwords"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/golovpeter/passbox_backend/internal/database"
)

func Register(db database.Database) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var in registerIn

		if err := ctx.BodyParser(&in); err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusUnprocessableEntity, 1, err.Error())
		}

		if in.Email == "" || in.Password == "" {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "Incorrect data input")
		}

		elementExist := false
		err := db.ExistUser(&elementExist, in.Email)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusInternalServerError, 1, err.Error())
		}

		if elementExist {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "User already registered!")
		}

		passwordHash := hash_passwords.GeneratePasswordHash(in.Password)

		_, err = db.InsertUser(in.Email, passwordHash)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusInternalServerError, 1, err.Error())
		}

		return make_response.MakeInfoResponse(ctx, fiber.StatusOK, 0, "Registration was successful!")
	}
}
