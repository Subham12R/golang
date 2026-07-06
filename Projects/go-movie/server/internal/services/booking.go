package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-movie-server/internal/store"
	"go-movie-server/internal/models"
	
)

type BookingRequest struct {
	ScreenID string `json:"screen_id" validate:"required"`
	SeatID	string `json:"seat_id" validate:"required"`
	ShowID	string `json:"show_id" validate:"required"`
	UserID	int `json:"user_id" validate:"required"`
}

type BookingService struct {
	store *store.SeatStore
	validate *validator.Validate
}

func NewBookingService(store *store.SeatStore) *BookingService {
    return &BookingService{
        store:    store,
        validate: validator.New(),
    }
}


func (s* BookingService) Book(req BookingRequest) (*models.Booking, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	_, err := s.store.GetScreenByID(req.ScreenID)
	if err != nil {
		return nil, fmt.Errorf("failed to book: %w", err)
	}

	booking, err := s.store.BookSeat(req.SeatID, req.ScreenID, req.ShowID, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to book: %w", err)
	}

	return booking, nil
}