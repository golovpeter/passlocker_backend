package configure_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/login"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/refresh_tokens"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/register"
	"github.com/golovpeter/passbox_backend/internal/api/middlewares/check_auth"
	"github.com/jmoiron/sqlx"
)

func ConfigureRouter(app *fiber.App, db *sqlx.DB) {
	//Middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${time} ${status} - ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.Config{
		//AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use("/api/p/", check_auth.CheckAuth(db))

	//Authentication endpoints
	app.Post("api/register", register.Register(db))
	app.Post("api/auth/login", login.Login(db))
	app.Post("api/refresh-tokens", refresh_tokens.RefreshTokens(db))
}
