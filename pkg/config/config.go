package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DevEnv = "dev"
)

type Config struct {
	App struct {
		Environment string
		Version     string
		Addr        string
		Port        string
	}
}

func New() *Config {
	var cfg Config

	if os.Getenv("APP_ENV") == DevEnv {
		err := godotenv.Load()
		if err != nil {
			return &cfg
		}
	}

	cfg.App.Environment = getEnv("APP_ENV", DevEnv)
	cfg.App.Version = getEnv("APP_VERSION", "0")
	cfg.App.Addr = getEnv("APP_ADDR", "localhost")
	cfg.App.Port = getEnv("APP_PORT", "8080")

	return &cfg
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf(
			"%s environment variable is not defined, so default value %s is used",
			key,
			defaultValue,
		)
		return defaultValue
	}
	return value
}
