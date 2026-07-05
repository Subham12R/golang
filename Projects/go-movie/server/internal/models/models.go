package models

import "time"

type Movie struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Genre    string  `json:"genre"`
	Duration int     `json:"duration"`
	Rating   float64 `json:"rating"`
}

type Show struct {
	ID        string    `json:"id"`
	MovieID   string    `json:"movie_id"`
	StartTime time.Time `json:"time"`
	TheaterID string    `json:"theater_id"`
}

type SeatStatus string

const (
	SeatAvailable   SeatStatus = "available"
	SeatBooked      SeatStatus = "booked"
	SeatReserved    SeatStatus = "reserved"
	SeatUnavailable SeatStatus = "unavailable"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Screen struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Seats map[string]Seat `json:"seats"`
}

type ShowSeat struct {
	ShowID string
	SeatID string
	Status SeatStatus
}

type Seat struct {
	Row    int        `json:"row"`
	Col    int        `json:"col"`
	SeatID string     `json:"seat_id"`
	Status SeatStatus `json:"status"`
}

type BookingStatus string

const (
	BookingPending   BookingStatus = "pending"
	BookingConfirmed BookingStatus = "confirmed"
	BookingCancelled BookingStatus = "cancelled"
)

type Booking struct {
	ID        string        `json:"id"`
	UserID    int           `json:"user_id"`
	ShowID    string        `json:"show_id"`
	ScreenID  string        `json:"screen_id"`
	SeatIDs   []string      `json:"seat_ids"`
	Status    BookingStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
}
