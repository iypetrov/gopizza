package database

import (
	"database/sql"
	"fmt"
	"github.com/iypetrov/gopizza/pkg/app"
	_ "github.com/lib/pq"
)

func New() *sql.DB {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		app.Cfg.Database.Username,
		app.Cfg.Database.Password,
		app.Cfg.Database.Host,
		app.Cfg.Database.Port,
		app.Cfg.Database.Name,
		app.Cfg.Database.SSL,
	)

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	return conn
}
