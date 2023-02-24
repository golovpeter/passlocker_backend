package configure_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/login"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/register"
	"github.com/golovpeter/passbox_backend/internal/api/middlewares/db_connect"
)

func ConfigureRouter(app *fiber.App) {
	// Middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${time} ${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.Config{
		//AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(db_connect.SetupDB())

	// Authentication endpoints
	app.Post("/register", register.Register)
	app.Post("/login", login.Login)
}
