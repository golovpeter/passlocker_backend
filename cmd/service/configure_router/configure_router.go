package configure_router

import (
	"1password_copy_project/internal/api/handlers/register"
	"1password_copy_project/internal/api/middlewares/db_connect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ConfigureRouter(app *fiber.App) {
	// Middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${time} ${status} - ${method} ${path}\n",
	}))
	app.Use(db_connect.SetupDB())

	// Authentication endpoints
	app.Post("/register", register.Register)
}
