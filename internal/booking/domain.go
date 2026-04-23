package booking

import (
	"errors"
	"time"
)

// Booking struct as per movie for seat reservation
type Booking struct {
	MoviedID  string
	UserID    string
	Status    string
	ID        string
	SeatID    string
	ExpiresAt time.Time
}

// custom errors
var (
	ErrSeatsAlreadyBooked = errors.New("Seat Already Booked!")
)

type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}
