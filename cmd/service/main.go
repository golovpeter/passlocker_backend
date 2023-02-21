package main

import (
	"log"
	"os"

	"1password_copy_project/cmd/service/configure_router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	app := fiber.New()
	configure_router.ConfigureRouter(app)

	if err := app.Listen(os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
