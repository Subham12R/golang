package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-movie-server/internal/handlers"
	"go-movie-server/internal/services"
	"go-movie-server/internal/store"
)

func main() {
	s := store.NewSeatStore()

	s.InitializeScreen("SCR-1", 8, 10)

	bookingService := services.NewBookingService(s)
	movieHandler := handlers.NewMovieHandler(s, bookingService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == "OPTIONS" {
				w.Write([]byte("OK"))
				return
			}
			h.ServeHTTP(w, r)
		})
	})
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
