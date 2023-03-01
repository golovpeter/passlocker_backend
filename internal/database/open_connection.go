package database

import (
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

func OpenConnection() *sqlx.DB {
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

	return db
}
