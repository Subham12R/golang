package store

import (
	"fmt"
	"go-movie-server/internal/models"
	"sync"
	"time"
)

type SeatStore struct {
	sync.RWMutex

	Movies   map[string]*models.Movie
	Shows    map[string][]*models.Show
	Screens  map[string]*models.Screen
	Bookings map[string]*models.Booking

	nextBookingID int
}

func NewSeatStore() *SeatStore {
	return &SeatStore{
		Movies:        make(map[string]*models.Movie),
		Shows:         make(map[string][]*models.Show),
		Screens:       make(map[string]*models.Screen),
		Bookings:      make(map[string]*models.Booking),
		nextBookingID: 1,
	}
}

// BookSeat books a seat with exclusive write lock to prevent double booking
func (s *SeatStore) BookSeat(screenID, seatID, showID string, userID int) (*models.Booking, error) {
	s.Lock()
	defer s.Unlock()

	screen, ok := s.Screens[screenID]
	if !ok {
		return nil, fmt.Errorf("screen %s not found", screenID)
	}

	seat, ok := screen.Seats[seatID]
	if !ok {
		return nil, fmt.Errorf("seat %s not found", seatID)
	}

	if seat.Status == models.SeatBooked || seat.Status == models.SeatReserved {
		return nil, fmt.Errorf("seat %s already booked or reserved", seatID)
	}

	booking := &models.Booking{
		ID:        fmt.Sprintf("BK-%d", s.nextBookingID),
		SeatIDs:   []string{seatID},
		ScreenID:  screenID,
		UserID:    userID,
		ShowID:    showID,
		Status:    models.BookingConfirmed,
		CreatedAt: time.Now(),
	}
	s.Bookings[booking.ID] = booking
	s.nextBookingID++

	updatedSeat := seat
	updatedSeat.Status = models.SeatBooked
	screen.Seats[seatID] = updatedSeat

	return booking, nil
}

func (s *SeatStore) GetAvailableSeats(screenID string) ([]*models.Seat, error) {
	s.RLock()
	defer s.RUnlock()

	screen, ok := s.Screens[screenID]
	if !ok {
		return nil, fmt.Errorf("screen %s not found", screenID)
	}

	var availableSeats []*models.Seat
	for _, seat := range screen.Seats {
		if seat.Status == models.SeatAvailable {
			s := seat 
			availableSeats = append(availableSeats, &s)
		}
	}

	return availableSeats, nil
}

func (s *SeatStore) InitializeScreen(screenID string, rows int, cols int) {
	s.Lock()
	defer s.Unlock()

	screen := &models.Screen{
		ID:    screenID,
		Name:  fmt.Sprintf("Screen %s", screenID),
		Seats: make(map[string]models.Seat),
	}

	for row := 1; row <= rows; row++ {
		for col := 1; col <= cols; col++ {
			seatID := fmt.Sprintf("%c%d", 'A'+(row-1), col)
			screen.Seats[seatID] = models.Seat{
				Row:    row,
				Col:    col,
				SeatID: seatID,
				Status: models.SeatAvailable,
			}
		}
	}

	s.Screens[screenID] = screen
}

func (s *SeatStore) GetScreenByID(screenID string) (*models.Screen, error) {
	s.RLock()
	defer s.RUnlock()

	screen, ok := s.Screens[screenID]
	if !ok {
		return nil, fmt.Errorf("screen not found")
	}

	return screen, nil
}

func (s *SeatStore) ListMovies() []*models.Movie {
	s.RLock()
	defer s.RUnlock()

	movies := make([]*models.Movie, 0, len(s.Movies))
	for _, movie := range s.Movies {
		movies = append(movies, movie)
	}

	return movies
}

func (s *SeatStore) GetMovieByID(movieID string) (*models.Movie, error) {
	s.RLock()
	defer s.RUnlock()

	movie, ok := s.Movies[movieID]
	if !ok {
		return nil, fmt.Errorf("movie %s not found", movieID)
	}

	return movie, nil
}

func (s *SeatStore) GetShowsForMovie(movieID string, date time.Time) []*models.Show {
	s.RLock()
	defer s.RUnlock()

	var shows []*models.Show
	for _, show := range s.Shows[movieID] {
		if sameDay(show.StartTime, date) {
			shows = append(shows, show)
		}
	}

	return shows
}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}

func (s *SeatStore) GetBookingByID(bookingID string) (*models.Booking, error) {
	s.RLock()
	defer s.RUnlock()

	booking, ok := s.Bookings[bookingID]
	if !ok {
		return nil, fmt.Errorf("booking %s not found", bookingID)
	}

	return booking, nil
}

func (s *SeatStore) CancelBooking(bookingID string) error {
	s.Lock()
	defer s.Unlock()

	booking, ok := s.Bookings[bookingID]
	if !ok {
		return fmt.Errorf("booking %s not found", bookingID)
	}

	if booking.Status == models.BookingCancelled {
		return fmt.Errorf("booking %s is already cancelled", bookingID)
	}

	screen, ok := s.Screens[booking.ScreenID]
	if ok {
		for _, seatID := range booking.SeatIDs {
			if seat, exists := screen.Seats[seatID]; exists {
				updatedSeat := seat
				updatedSeat.Status = models.SeatAvailable
				screen.Seats[seatID] = updatedSeat
			}
		}
	}

	booking.Status = models.BookingCancelled
	return nil
}
