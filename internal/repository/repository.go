package repository

import (
	"database/sql"
)

type Weather interface {
	UploadWeather(city string) error
}

type Repository struct {
	Weather
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
