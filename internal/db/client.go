package db

import (
	"database/sql"
	"fmt"

	"github.com/ctuzelov/weather-api/internal/config"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewClient(cfg *config.Config) (*ClientImpl, error) {
	pgURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.PgUser, cfg.PgPassword, cfg.PgHost, cfg.PgPort, cfg.PgName, cfg.PgSSLMode)

	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		cfg.L.Error("cannot connect to db", zap.Error(err))
		return nil, fmt.Errorf("cannot connect to db, %w", err)
	}

	cfg.L.Info("connected to db")

	return &ClientImpl{db: db}, nil
}

type ClientImpl struct {
	db *sql.DB
}

func (db *ClientImpl) GetDB() *sql.DB {
	return db.db
}

// func runDBMigration(sourceURL, databaseURL string) error {
// 	migration, err := migrate.New(sourceURL, databaseURL)
// 	if err != nil {
// 		return fmt.Errorf("cannot create migration, %w", err)
// 	}

// 	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
// 		return fmt.Errorf("cannot up migration, %w", err)
// 	}

// 	return nil
// }
