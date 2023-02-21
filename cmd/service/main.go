package main

import (
	"1password_copy_project/cmd/service/configure_router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

func main() {
	//db, err := sqlx.Connect("pgx", os.Getenv("POSTGRESQL_URL"))
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(5)
	//db.SetConnMaxLifetime(5 * time.Minute)
	//db.SetConnMaxIdleTime(5 * time.Minute)

	//defer db.Close()

	app := fiber.New()
	configure_router.ConfigureRouter(app)

	if err := app.Listen(os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
