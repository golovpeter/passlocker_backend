package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/cmd/service/configure_router"
	"github.com/golovpeter/passbox_backend/internal/database"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	app := fiber.New()

	db := database.OpenConnection()
	defer db.Close()

	configure_router.ConfigureRouter(app, db)

	if err := app.Listen(os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
