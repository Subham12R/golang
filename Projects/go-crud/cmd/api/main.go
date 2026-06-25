package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/subham12r/go-crud/internal/database"
	"github.com/subham12r/go-crud/internal/handlers"
	"github.com/subham12r/go-crud/internal/repository"
	"github.com/subham12r/go-crud/internal/service"
)

func main(){
	if err := godotenv.Load();
	err != nil {
		log.Println("No .env file found")
	}
	

	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	movieRepo := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepo)
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
	log.Println("Server Started running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)


}