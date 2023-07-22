package config

import "os"

type Config struct {
	DbPath string
	JwtKey string
	Port   string
}

func NewConfig() *Config {
	return &Config{
		DbPath: os.Getenv("POSTGRESQL_URL"),
		JwtKey: os.Getenv("JWT_KEY"),
		Port:   os.Getenv("PORT"),
	}
}
