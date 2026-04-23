package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBURL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return &Config{}, err
	}

	slog.Info("env loaded successfully")

	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBURL:   os.Getenv("DB_URL"),
	}, nil
}
