package delete_password

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/internal/common/auth_tokens"
	"github.com/golovpeter/passbox_backend/internal/common/make_response"
	"github.com/golovpeter/passbox_backend/internal/common/parse_headers"
	"github.com/golovpeter/passbox_backend/internal/database"
	"strconv"
)

func DeletePassword(db database.Passwords) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		accessToken, err := parse_headers.ParseAuthHeader(ctx)

		passwordID, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "invalid id")
		}

		claims, err := auth_tokens.GetTokenClaims(accessToken)

		if err != nil {
			return err
		}

		var passwordUserId int
		err = db.SelectPasswordUserID(&passwordUserId, passwordID)

		if passwordUserId == 0 {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, "there is no such password")
		}

		if passwordUserId != int(claims["UserID"].(float64)) {
			return make_response.MakeInfoResponse(
				ctx,
				fiber.StatusBadRequest,
				1,
				"this password belongs to another user",
			)
		}

		_, err = db.DeletePassword(passwordID)

		if err != nil {
			return make_response.MakeInfoResponse(ctx, fiber.StatusBadRequest, 1, err.Error())
		}

		return make_response.MakeInfoResponse(ctx, fiber.StatusOK, 0, "password successful deleted")
	}
}
