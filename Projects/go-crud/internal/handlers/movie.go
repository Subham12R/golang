package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/subham12r/go-crud/internal/models"
	"github.com/subham12r/go-crud/internal/service"
)

// var movies = []models.Movie{
// 	{
// 		ID: 		1,
// 		Title:  	"Interstellar",
// 		Director: 	"Christopher Nolan",
// 		ReleaseYear: 2014,
// 	},
// 	{
// 		ID:			2,
// 		Title: 		"Odyssey",
// 		Director: 	"Christopher Nolan",
// 		ReleaseYear: 2026,
// 	},
// }

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(
	service *service.MovieService,
) *MovieHandler {
	return &MovieHandler{
		service: service,
	}
}


func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request){
	movies, err:= h.service.GetMovies()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	
	json.NewEncoder(w).Encode(movies)
}


func (h* MovieHandler) GetByID(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusNotFound)
		return
	}

	movie, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "Movie not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func(h *MovieHandler) Create(w http.ResponseWriter, r *http.Request){
	var req models.CreateMovieRequest

	if err := json.NewDecoder(r.Body).Decode(&req); 
	err != nil {
		http.Error(w, "Invalid JSON Format", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Invalid Title", http.StatusBadRequest)
		return
	}
	movie, err := h.service.Create(req)
	if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}


func(h *MovieHandler) Update(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}

	var req models.UpdateMovieRequest
	if err := json.NewDecoder(r.Body).Decode(&req);
	err != nil {
		http.Error(w, "Invalid JSON Format", http.StatusBadRequest)
		return
	}

	movie, err := h.service.Update(id, req)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func(h *MovieHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}

	if err := h.service.Delete(id);
	err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
// func GetMovies(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)
// 	json.NewEncoder(w).Encode(movies)
// }


// func GetMovie(w http.ResponseWriter, r *http.Request){
// 	idStr := r.PathValue("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid id", http.StatusBadRequest)
// 		return
// 	}

// 	for _, movie := range movies {
// 		if movie.ID == id {
// 			w.Header().Set(
// 				"Content-Type",
// 				"application/json",
// 			)
// 			json.NewEncoder(w).Encode(movie)
// 			return
// 		}	
// 	}
// 	http.Error(
// 		w, "Movie Not Found", http.StatusNotFound,
// 	)
// }
// func CreateMovie(w http.ResponseWriter, r *http.Request) {
// 	var req models.CreateMovieRequest

// 	err := json.NewDecoder(r.Body).Decode(&req)

// 	if err != nil {
// 		http.Error(w, "Invalid JSON Format", http.StatusBadRequest,)
// 		return
// 	}
		
// 	if req.Title == "" {
// 		http.Error(w, "Title is required", http.StatusBadRequest)
// 		return
// 	}

// 	movie := models.Movie{
// 		ID: 			len(movies) + 1,
// 		Title: 			req.Title,
// 		Director: 		req.Director,
// 		ReleaseYear:	req.ReleaseYear,
// 	}

// 	movies = append(movies, movie)

// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)
// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(movie)
// }

// func UpdateMovie(w http.ResponseWriter, r *http.Request){
// 	idStr := r.PathValue("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil{
// 		http.Error(w, "Invalid id", http.StatusBadRequest)
// 		return
// 	}

// 	var req models.UpdateMovieRequest

// 	err = json.NewDecoder(r.Body).Decode(&req)
	
// 	if err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	for i := range movies{
// 		if movies[i].ID == id{
// 			movies[i].Title = req.Title
// 			movies[i].Director = req.Director
// 			movies[i].ReleaseYear = req.ReleaseYear

// 			w.Header().Set(
// 				"Content-Type",
// 				"application/json",
// 			)

// 			json.NewEncoder(w).Encode(movies[i])
			
// 			return
// 		}	
// 	}
	
// 	http.Error(w, "Movie not found", http.StatusNotFound)
// }

// func DeleteMovie(w http.ResponseWriter, r *http.Request){
// 	idStr := r.PathValue("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil{
// 		http.Error(w, "Invalid id", http.StatusBadRequest)
// 		return
// 	}

// 	for i := range movies{
// 		if movies[i].ID == id{
// 			movies = append(
// 				movies[:i],
// 				movies[i+1:]...,
// 			)
			
// 			w.WriteHeader(http.StatusNoContent)
			
// 		return
// 		}
// 	}
// 	http.Error(w, "Movie not found", http.StatusNotFound)
// }