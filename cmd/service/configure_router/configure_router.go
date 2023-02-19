package configure_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
)

func ConfigureRouter(app *fiber.App, conn *sqlx.DB) {
	// Middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals("dbConn", conn)
		return ctx.Next()
	})
}
