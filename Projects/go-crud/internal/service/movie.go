package service

import (
	"errors"
	"github.com/subham12r/go-crud/internal/models"
)

var ErrNotFound = errors.New("Movie not found")
type MovieService struct {
	movies []models.Movie
	nextID int
}

func NewMovieService() *MovieService {
	return &MovieService{
		movies: []models.Movie{
			{
				ID:          1,
				Title:       "Interstellar",
				Director:    "Christopher Nolan",
				ReleaseYear: 2014,
			},
			{
				ID:          2,
				Title:       "Odyssey",
				Director:    "Christopher Nolan",
				ReleaseYear: 2026,
			},
		},
		nextID: 3,
	}
}

func (s *MovieService) GetMovies() []models.Movie {
	return s.movies
}

func (s *MovieService) GetByID(id int) (models.Movie, error){
	for _, movie := range s.movies {
		if movie.ID == id {
			return movie, nil
		}
	}
	return models.Movie{}, ErrNotFound
} 

func (s *MovieService) Create(req models.CreateMovieRequest) models.Movie{
	movie := models.Movie{
		ID: 			s.nextID,
		Title:  		req.Title,
		Director: 		req.Director,
		ReleaseYear:	req.ReleaseYear,
	}
	s.nextID++
	s.movies = append(s.movies, movie)
	return movie
}

func (s *MovieService) Update(id int, req models.UpdateMovieRequest) (models.Movie, error) {
	for i:= range s.movies{
		if s.movies[i].ID == id {
			s.movies[i].Title = req.Title
			s.movies[i].Director = req.Director
			s.movies[i].ReleaseYear = req.ReleaseYear
			return s.movies[i], nil
		}
	}
	return models.Movie{}, ErrNotFound
}

func (s *MovieService) Delete(id int) error {
	for i:= range s.movies{
		if s.movies[i].ID == id {
			s.movies = append(s.movies[:i], s.movies[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}