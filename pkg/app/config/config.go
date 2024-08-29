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

	Database struct {
		Name     string
		Username string
		Password string
		Host     string
		Port     string
		SSL      string
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
	cfg.Database.Name = getEnv("APP_DB_NAME", "goshop")
	cfg.Database.Username = getEnv("APP_DB_USERNAME", "user")
	cfg.Database.Password = getEnv("APP_DB_PASSWORD", "pass")
	cfg.Database.Host = getEnv("APP_DB_HOST", "localhost")
	cfg.Database.Port = getEnv("APP_DB_PORT", "5432")
	cfg.Database.SSL = getEnv("APP_DB_SSL", "disable")

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
