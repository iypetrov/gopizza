package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DevLocal = "local"
	cfg      *Config
)

type Config struct {
	App struct {
		Environment string
		Version     string
		Addr        string
		Port        string
		Secret      string
	}

	Database struct {
		Name     string
		Username string
		Password string
		Host     string
		Port     string
		SSL      string
	}

	AWS struct {
		Region          string
		AccessKeyID     string
		SecretAcessKey  string
		S3BucketName    string
		CognitoClientID string
	}
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cfg = &Config{}
	cfg.App.Environment = getEnv("APP_ENV", DevLocal)
	cfg.App.Version = getEnv("APP_VERSION", "0")
	cfg.App.Addr = getEnv("APP_ADDR", "localhost")
	cfg.App.Port = getEnv("APP_PORT", "8080")
	cfg.App.Secret = getEnv("APP_SECRET", "default")
	cfg.Database.Name = getEnv("APP_DB_NAME", "gopizza")
	cfg.Database.Username = getEnv("APP_DB_USERNAME", "user")
	cfg.Database.Password = getEnv("APP_DB_PASSWORD", "pass")
	cfg.Database.Host = getEnv("APP_DB_HOST", "localhost")
	cfg.Database.Port = getEnv("APP_DB_PORT", "5432")
	cfg.Database.SSL = getEnv("APP_DB_SSL", "disable")
	cfg.AWS.Region = getEnv("AWS_REGION", "default")
	cfg.AWS.AccessKeyID = getEnv("AWS_ACCESS_KEY_ID", "default")
	cfg.AWS.SecretAcessKey = getEnv("AWS_SECRET_ACCESS_KEY", "default")
	cfg.AWS.S3BucketName = getEnv("S3_BUCKET_NAME", "default")
	cfg.AWS.CognitoClientID = getEnv("COGNITO_CLIENT_ID", "default")
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
	basePath := c.App.Addr

	if c.App.Environment == DevLocal {
		protocol = "http://"
		basePath = fmt.Sprintf("%s:%s", c.App.Addr, c.App.Port)
	}

	return fmt.Sprintf("%s%s", protocol, basePath)
}

func (c *Config) GetAdminPrefix() string {
	return "/admin"
}

func (c *Config) GetClientAPIPrefix() string {
	return fmt.Sprintf("/api/v%s", c.App.Version)
}

func (c *Config) GetAdminAPIPrefix() string {
	return fmt.Sprintf("/admin/v%s", c.App.Version)
}
