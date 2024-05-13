package repository

import "github.com/msaufi2325/go-movies-back-end/internal/models"

type DatabaseRepo interface {
	AllMovies() ([]*models.Movie, error)
}
