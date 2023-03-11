package log_out

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/golovpeter/passbox_backend/internal/common/parse_headers"
	"github.com/jmoiron/sqlx"
)

func LogOut(conn *sqlx.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		accessToken, _ := parse_headers.ParseAuthHeader(ctx)

		_, err := conn.Exec("delete from tokens where access_token = $1", accessToken)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		return make_response.MakeInfoResponse(ctx, fiber.StatusOK, 0, "token successful deleted")
	}
}
