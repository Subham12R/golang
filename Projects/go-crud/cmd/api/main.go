package main

import (
	"log"
	"net/http"
	"github.com/subham12r/go-crud/internal/handlers"
	"github.com/subham12r/go-crud/internal/service"
)

func main(){
		
	mux := http.NewServeMux()
	movieService := service.NewMovieService()
	movieHandler := handlers.NewMovieHandler(movieService)	

	mux.HandleFunc(
		"GET /health",
		func(w http.ResponseWriter, r *http.Request){
			w.Write([]byte("OK"))
		},
	)

	mux.HandleFunc(
		"GET /movies",
		movieHandler.GetMovies,
	)

	mux.HandleFunc(
		"GET /movies/{id}",
		movieHandler.GetByID,
	)

	mux.HandleFunc(
		"POST /movies",
		movieHandler.Create,
	)

	mux.HandleFunc(
		"PUT /movies/{id}",
		movieHandler.Update,
	)
	mux.HandleFunc(
		"DELETE /movies/{id}",
		movieHandler.Delete,
	)
	log.Println("Server Started")
	http.ListenAndServe(":8080", mux)


}