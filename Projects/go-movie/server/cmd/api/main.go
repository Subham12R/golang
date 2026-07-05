package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-movie-server/internal/handlers"
	"go-movie-server/internal/store"
)

func main() {
	s := store.NewSeatStore()

	s.InitializeScreen("SCR-1", 8, 10)

	movieHandler := handlers.NewMovieHandler(s)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Movie Booking API - Healthy"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/movies", movieHandler.GetMovies)
		r.Get("/seats", movieHandler.GetAvailableSeats)	
		r.Post("/bookings", movieHandler.HandleBookSeat)
		r.Get("/bookings/{bookingID}", movieHandler.HandleGetBooking)
		r.Delete("/bookings/{bookingID}", movieHandler.HandleCancelBooking)
	})

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
