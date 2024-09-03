package config

import (
	"database/sql"
	"fmt"
	"github.com/iypetrov/gopizza/internal/database"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func CreateDatabaseConnection(cfg *Config) (*sql.DB, error) {
	url := getDataBaseUrl(cfg)
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewDatabase(conn *sql.DB) *database.Queries {
	return database.New(conn)
}

func RunSchemaMigration(conn *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(conn, "sql/migrations"); err != nil {
		return err
	}

	return nil
}

func getDataBaseUrl(cfg *Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSL,
	)
}
