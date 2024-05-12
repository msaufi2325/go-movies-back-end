package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/msaufi2325/go-movies-back-end/internal/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "1994-10-14")

	shawshank := models.Movie{
		ID:          1,
		Title:       "The Shawshank Redemption",
		ReleaseDate: rd,
		Runtime:     142,
		MPAARating:  "R",
		Description: "Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, shawshank)

	rd, _ = time.Parse("2006-01-02", "1972-03-24")

	godfather := models.Movie{
		ID:          2,
		Title:       "The Godfather",
		ReleaseDate: rd,
		Runtime:     175,
		MPAARating:  "R",
		Description: "An organized crime dynasty's aging patriarch transfers control of his clandestine empire to his reluctant son.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, godfather)

	rd, _ = time.Parse("2006-01-02", "2008-07-18")
	darkKnight := models.Movie{
		ID:          3,
		Title:       "The Dark Knight",
		ReleaseDate: rd,
		Runtime:     152,
		MPAARating:  "PG-13",
		Description: "When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, darkKnight)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
