//go:build prod 
// +build prod

package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var cfg *Config

type Config struct {
	App struct {
		Profile string
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

	AWS struct {
		Region          string
		AccessKeyID     string
		SecretAcessKey  string
		S3BucketName    string
		CognitoClientID string
	}

	Stripe struct {
		PublishableKey string
		SecretKey      string
		WebhookSecret  string
	}
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cfg = &Config{}
	cfg.App.Profile = string(Prod) 
	cfg.App.Version = os.Getenv("APP_VERSION")
	cfg.App.Addr = os.Getenv("APP_ADDR")
	cfg.App.Port = os.Getenv("APP_PORT")
	cfg.Database.Name = os.Getenv("APP_DB_NAME")
	cfg.Database.Username = os.Getenv("APP_DB_USERNAME")
	cfg.Database.Password = os.Getenv("APP_DB_PASSWORD")
	cfg.Database.Host = os.Getenv("APP_DB_HOST")
	cfg.Database.Port = os.Getenv("APP_DB_PORT")
	cfg.Database.SSL = os.Getenv("APP_DB_SSL")
	cfg.AWS.Region = os.Getenv("AWS_REGION")
	cfg.AWS.AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	cfg.AWS.SecretAcessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	cfg.AWS.S3BucketName = os.Getenv("S3_BUCKET_NAME")
	cfg.AWS.CognitoClientID = os.Getenv("COGNITO_CLIENT_ID")
	cfg.Stripe.PublishableKey = os.Getenv("STRIPE_PUBLISHABLE_KEY")
	cfg.Stripe.SecretKey = os.Getenv("STRIPE_SECRET_KEY")
	cfg.Stripe.WebhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")
}

func Get() *Config {
	return cfg
}

func (c *Config) GetBaseWebUrl() string {
	protocol := "https://"
	basePath := c.App.Addr

	if c.App.Profile == string(Local) {
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
