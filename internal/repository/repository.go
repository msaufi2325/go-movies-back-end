package repository

import (
	"database/sql"

	"github.com/msaufi2325/go-movies-back-end/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
