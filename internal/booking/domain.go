package booking

import (
	"context"
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
	Book(b Booking) (Booking, error)
	ListBookings(movieID string) []Booking
	Confirm(ctx context.Context, sessionID string, userID string) (Booking, error)
	Release(ctx context.Context, sessionID string, userID string) error
}
