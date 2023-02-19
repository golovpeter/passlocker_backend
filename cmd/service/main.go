package main

import (
	"1password_copy_project/cmd/service/configure_router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func main() {
	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	app := fiber.New()
	configure_router.ConfigureRouter(app, db)

	if err = app.Listen(os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
