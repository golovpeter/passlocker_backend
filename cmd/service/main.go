package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golovpeter/passbox_backend/cmd/service/configure_router"
	"github.com/golovpeter/passbox_backend/internal/database/postgresql"
)

func main() {
	app := fiber.New()

	db := postgresql.NewDatabase()
	defer db.Conn.Close()

	configure_router.ConfigureRouter(app, db)

	app.Static("/", "../../static")
	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("../../static/index.html")
	})

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
