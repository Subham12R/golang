package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/subham12r/go-crud/internal/models"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func(r *MovieRepository) GetMovies() ([]models.Movie, error){
	rows, err := r.db.Query("SELECT * FROM movies ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("querying movies: %w", err)
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		if err := rows.Scan(&m.ID, &m.Title, &m.Director, &m.ReleaseYear);
		err != nil {
			return nil, fmt.Errorf("scanning movie: %w", err)
		}
		movies = append(movies, m)
	}
	return movies, rows.Err()
}

func(r *MovieRepository) GetByID(id int) (models.Movie, error){
	var m models.Movie
	err := r.db.QueryRow("SELECT * FROM movies WHERE id = $1", id,).Scan(&m.ID, &m.Title, &m.Director, &m.ReleaseYear) 
	if errors.Is(err, sql.ErrNoRows) {
		return models.Movie{}, models.ErrNotFound
	}
	
	if err != nil {
		return models.Movie{}, fmt.Errorf("Querying movie %d: %w", id, err) 
	}
	return m, nil
}

func (r *MovieRepository) Create(req models.CreateMovieRequest) (models.Movie, error) {
    var m models.Movie
    err := r.db.QueryRow(
        `INSERT INTO movies (title, director, release_year)
         VALUES ($1, $2, $3)
         RETURNING id, title, director, release_year`,
        req.Title, req.Director, req.ReleaseYear,
    ).Scan(&m.ID, &m.Title, &m.Director, &m.ReleaseYear)

    if err != nil {
        return models.Movie{}, fmt.Errorf("inserting movie: %w", err)
    }
    return m, nil
}

func (r *MovieRepository) Update(id int, req models.UpdateMovieRequest) (models.Movie, error) {
    var m models.Movie
    err := r.db.QueryRow(
        `UPDATE movies SET title=$1, director=$2, release_year=$3
         WHERE id=$4
         RETURNING id, title, director, release_year`,
        req.Title, req.Director, req.ReleaseYear, id,
    ).Scan(&m.ID, &m.Title, &m.Director, &m.ReleaseYear)

    if errors.Is(err, sql.ErrNoRows) {
        return models.Movie{}, models.ErrNotFound
    }
    if err != nil {
        return models.Movie{}, fmt.Errorf("updating movie %d: %w", id, err)
    }
    return m, nil
}

func (r *MovieRepository) Delete(id int) error {
    result, err := r.db.Exec(
        "DELETE FROM movies WHERE id = $1", id,
    )
    if err != nil {
        return fmt.Errorf("deleting movie %d: %w", id, err)
    }

    n, _ := result.RowsAffected()
    if n == 0 {
        return models.ErrNotFound
	}
    return nil
}
