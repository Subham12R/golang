package handlers

import (
	"encoding/json"
	"go-movie-server/internal/services"
	"go-movie-server/internal/store"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type MovieHandler struct {
	store          *store.SeatStore
	bookingService *services.BookingService
}

func NewMovieHandler(store *store.SeatStore, bookingService *services.BookingService) *MovieHandler {
	return &MovieHandler{store: store, bookingService: bookingService}
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies := h.store.ListMovies()

	response, _ := json.MarshalIndent(movies, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *MovieHandler) GetShowtimes(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	dateStr := r.URL.Query().Get("date")
	if movieID == "" || dateStr == "" {
		http.Error(w, `{"error": "movieID and date are required"}`, http.StatusBadRequest)
		return
	}

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, `{"error": "invalid date"}`, http.StatusBadRequest)
			return
		}
	}

	shows := h.store.GetShowsForMovie(movieID, date)

	type showResponse struct {
		ID     string `json:"id"`
		Time   string `json:"time"`
		Screen string `json:"screen"`
	}

	resp := make([]showResponse, 0, len(shows))
	for _, show := range shows {
		resp = append(resp, showResponse{
			ID:     show.ID,
			Time:   show.StartTime.Format("3:04 PM"),
			Screen: show.ScreenID,
		})
	}

	response, _ := json.MarshalIndent(resp, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *MovieHandler) GetAvailableSeats(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	screenID := query.Get("screen_id")
	if screenID == "" {
		http.Error(w, `{"error": "screen_id is required"}`, http.StatusBadRequest)
		return
	}

	seats, err := h.store.GetAvailableSeats(screenID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response, _ := json.MarshalIndent(seats, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *MovieHandler) HandleBookSeat(w http.ResponseWriter, r *http.Request) {
	var req services.BookingRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	booking, err := h.bookingService.Book(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	response, _ := json.MarshalIndent(booking, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *MovieHandler) HandleGetBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := r.URL.Query().Get("booking_id")
	if bookingID == "" {
		http.Error(w, `{"error": "booking_id is required"}`, http.StatusBadRequest)
		return
	}

	booking, err := h.store.GetBookingByID(bookingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response, _ := json.MarshalIndent(booking, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *MovieHandler) HandleCancelBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := r.URL.Query().Get("booking_id")
	if bookingID == "" {
		http.Error(w, `{"error": "booking_id is required"}`, http.StatusBadRequest)
		return
	}

	err := h.store.CancelBooking(bookingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "booking cancelled successfully"}`))
}
