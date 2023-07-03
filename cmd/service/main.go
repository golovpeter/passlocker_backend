package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golovpeter/passbox_backend/cmd/service/configure_router"
	"github.com/golovpeter/passbox_backend/internal/config"
	"github.com/golovpeter/passbox_backend/internal/database/postgresql"
)

func main() {
	app := fiber.New()

	config := config.NewConfig()

	db := postgresql.NewDatabase(config.DbPath)
	defer db.Conn.Close()

	logs := logger.New(logger.Config{
		Format: "[${ip}]:${port} ${time} ${status} - ${method} ${path}\n",
	})

	log.Println("\n" +
		"   ___               __          __          " + "\n" +
		"  / _ \\___ ____ ___ / / ___ ____/ /_____ ____" + "\n" +
		" / ___/ _ `(_-<(_-</ /_/ _ / __/  '_/ -_/ __/" + "\n" +
		"/_/   \\_,_/___/___/____\\___\\__/_/\\_\\\\__/_/   " + "\n" +
		"                                             ")

	configure_router.ConfigureRouter(app, db, logs, config)

	app.Static("/", "../../static")
	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("../../static/index.html")
	})

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
