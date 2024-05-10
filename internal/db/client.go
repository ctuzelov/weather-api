package db

import (
	"fmt"

	"github.com/ctuzelov/weather-api/internal/config"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewClient(cfg *config.Config) (*ClientImpl, error) {
	pgURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.PgUser, cfg.PgPassword, cfg.PgHost, cfg.PgPort, cfg.PgName, cfg.PgSSLMode)

	db, err := sqlx.Connect("postgres", pgURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("connected successfully")

	return &ClientImpl{db: db}, nil
}

type ClientImpl struct {
	db *sqlx.DB
}

func (db *ClientImpl) GetDB() *sqlx.DB {
	return db.db
}
