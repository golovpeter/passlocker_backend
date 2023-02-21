package make_response

import "github.com/gofiber/fiber/v2"

func MakeInfoResponse(ctx *fiber.Ctx, httpStatus int, errorCode int, message string) error {
	return ctx.Status(httpStatus).JSON(fiber.Map{
		"errorCode": errorCode,
		"message":   message,
	})
}
