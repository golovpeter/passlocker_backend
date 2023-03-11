package configure_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/auth/log_out"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/auth/login"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/auth/refresh_tokens"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/auth/register"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/passwords/add_password"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/passwords/delete_password"
	"github.com/golovpeter/passbox_backend/internal/api/handlers/passwords/get_all_passwords"
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

	app.Use("api/p/", check_auth.CheckAuth(db))

	//Authentication endpoints
	app.Post("api/register", register.Register(db))
	app.Post("api/auth/login", login.Login(db))
	app.Post("api/refresh-tokens", refresh_tokens.RefreshTokens(db))
	app.Delete("api/log-out", log_out.LogOut(db))

	//Private endpoints
	app.Post("api/p/add-password", add_password.AddPassword(db))
	app.Post("api/p/delete-password", delete_password.DeletePassword(db))
	app.Get("api/p/get-all-passwords", get_all_passwords.GetAllNotes(db))
}
