package main

import (
	"log"
	"net/http"
	"github.com/subham12r/go-crud/internal/handlers"
)

func main(){
	
	mux := http.NewServeMux()
	
	mux.HandleFunc(
		"GET /health",
		func(w http.ResponseWriter, r *http.Request){
			w.Write([]byte("OK"))
		},
	)

	mux.HandleFunc(
		"GET /movies",
		handlers.GetMovies,
	)

	mux.HandleFunc(
		"GET /movies/{id}",
		handlers.GetMovie,
	)

	mux.HandleFunc(
		"POST /movies",
		handlers.CreateMovie,
	)

	mux.HandleFunc(
		"PUT /movies/{id}",
		handlers.UpdateMovie,
	)
	mux.HandleFunc(
		"DELETE /movies/{id}",
		handlers.DeleteMovie,
	)
	log.Println("Server Started")
	http.ListenAndServe(":8080", mux)


}