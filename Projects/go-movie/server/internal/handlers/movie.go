package handlers

import (
	"encoding/json"
	"go-movie-server/internal/store"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	store *store.SeatStore
}

func NewMovieHandler(store *store.SeatStore) *MovieHandler {
	return &MovieHandler{store: store}
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies := []struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Genre    string `json:"genre"`
		Duration int    `json:"duration"`
	}{
		{"1", "Inception", "Sci-Fi", 148},
		{"2", "The Matrix", "Action", 136},
		{"3", "Interstellar", "Sci-Fi", 169},
	}

	response, _ := json.MarshalIndent(movies, "", "  ")
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
	query := r.URL.Query()
	screenID := query.Get("screen_id")
	seatID := query.Get("seat_id")
	// Parse userID from query param (MVP: no auth middleware yet)
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, `{"error": "valid user_id is required"}`, http.StatusBadRequest)
		return
	}

	if screenID == "" || seatID == "" {
		http.Error(w, `{"error": "screen_id and seat_id are required"}`, http.StatusBadRequest)
		return
	}

	booking, err := h.store.BookSeat(screenID, seatID, r.URL.Query().Get("show_id"), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
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
