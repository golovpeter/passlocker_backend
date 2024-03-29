package get_all_passwords

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/internal/common/auth_tokens"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/golovpeter/passbox_backend/internal/common/parse_headers"
	"github.com/golovpeter/passbox_backend/internal/database"
)

func GetAllNotes(db database.Database) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		accessToken, err := parse_headers.ParseAuthHeader(ctx)
		claims, err := auth_tokens.GetTokenClaims(accessToken)

		if err != nil {
			return err
		}

		passwords := make([]PasswordsOut, 0)

		err = db.SelectAllPasswords(&passwords, int(claims["UserID"].(float64)))

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		return ctx.JSON(fiber.Map{
			"passwords": passwords,
		})
	}
}
