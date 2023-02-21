package register

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Register(ctx *fiber.Ctx) error {
	conn := ctx.Locals("dbConn").(*sqlx.DB)

	var in RegisterIn

	if err := ctx.BodyParser(&in); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errorCode": 1,
			"message":   err.Error(),
		})
	}

	if in.Email == "" || in.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorCode": 1,
			"message":   "Incorrect data input",
		})
	}

	elementExist := false
	err := conn.Get(&elementExist, "select exists(select email from users where email = $1)", in.Email)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errorCode": 1,
			"message":   err.Error(),
		})
	}

	if elementExist {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorCode": 1,
			"message":   "User already registered!",
		})
	}

	_, err = conn.Exec("insert into users (email, password) values ($1, $2)", in.Email, in.Password)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errorCode": 1,
			"message":   err.Error(),
		})
	}

	return nil
}
