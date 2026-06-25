package service

import "github.com/subham12r/go-crud/internal/models"

type MovieService struct{
	movies []models.Movie
	nextID int
}

func NewMovieService() *MovieService {
	return &MovieService{
		movies: []models.Movie{
			{
				ID: 		1,
				Title:  	"Interstellar",
				Director: 	"Christopher Nolan",
				ReleaseYear: 2014,
			},
			{
				ID:			2,
				Title: 		"Odyssey",
				Director: 	"Christopher Nolan",
				ReleaseYear: 2026,
			},
		},
		nextID: 3,
	}
}

func (s *MovieService) GetMovies() []models.Movie {
	return s.movies
}

