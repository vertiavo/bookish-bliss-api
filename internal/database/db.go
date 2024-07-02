package database

import (
	"database/sql"

	_ "github.com/lib/pq" // Postgres driver
	"github.com/vertiavo/bookish-bliss-api/internal/config"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
