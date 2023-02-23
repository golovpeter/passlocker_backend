package db_connect

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupDB() func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		db, err := sqlx.Connect("pgx", os.Getenv("POSTGRESQL_URL"))
		if err != nil {
			panic(err)
		}

		if err = db.Ping(); err != nil {
			panic(err)
		}

		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(5 * time.Minute)
		db.SetConnMaxIdleTime(5 * time.Minute)

		defer db.Close()

		ctx.Locals("dbConn", db)
		return ctx.Next()
	}
}
