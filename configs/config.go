package configs

import (
	"fmt"
	"log"
	"os"
)

var (
	DevEnv = "dev"
	cfg    *Config
)

type Config struct {
	App struct {
		Environment string
		Version     string
		Addr        string
		Port        string
	}
}

func Init() {
	cfg = &Config{}
	cfg.App.Environment = getEnv("APP_ENV", DevEnv)
	cfg.App.Version = getEnv("APP_VERSION", "0")
	cfg.App.Addr = getEnv("APP_ADDR", "localhost")
	cfg.App.Port = getEnv("APP_PORT", "8080")
}

func Get() *Config {
	return cfg
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("%s environment variable is not defined, so default value %s is used", key, defaultValue)
		return defaultValue
	}
	return value
}

func (c *Config) GetBaseWebUrl() string {
	protocol := "https://"
	basePath := fmt.Sprintf("%s", c.App.Addr)

	if c.App.Environment == DevEnv {
		protocol = "http://"
		basePath = fmt.Sprintf("%s:%s", c.App.Addr, c.App.Port)
	}

	return fmt.Sprintf("%s%s", protocol, basePath)
}

func (c *Config) GetAPIPrefix() string {
	return fmt.Sprintf("/api/v%s", c.App.Version)
}

func (c *Config) GetAdminPrefix() string {
	return "/admin"
}
