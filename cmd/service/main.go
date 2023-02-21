package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"

	"github.com/golovpeter/passbox_backend/cmd/service/configure_router"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	app := fiber.New()
	configure_router.ConfigureRouter(app)

	if err := app.Listen(os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
