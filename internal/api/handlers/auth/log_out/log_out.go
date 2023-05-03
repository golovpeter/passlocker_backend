package log_out

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/golovpeter/passbox_backend/internal/common/parse_headers"
	"github.com/golovpeter/passbox_backend/internal/database"
)

func LogOut(db database.Database) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		accessToken, _ := parse_headers.ParseAuthHeader(ctx)

		_, err := db.DeleteUser(accessToken)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		return make_response.MakeInfoResponse(ctx, fiber.StatusOK, 0, "token successful deleted")
	}
}
