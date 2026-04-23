package booking

import "context"

type Service struct {
	store BookingStore
}

func NewService(store BookingStore) *Service {
	return &Service{store}
}

func (s *Service) Book(b Booking) (Booking, error) {
	err := s.store.Book(b)
	if err != nil {
		// On failure, return an empty Booking struct and the error.
		return Booking{}, err
	}
	// On success, return the original Booking struct and a nil error.
	return b, nil
}

func (s *Service) ListBookings(movieID string) []Booking {
	return s.store.ListBookings(movieID)
}

func (s *Service) ConfirmSeat(ctx context.Context, sessionID string, userID string) (Booking, error) {
	return s.store.Confirm(ctx, sessionID, userID)
}

func (s *Service) ReleaseSeat(ctx context.Context, sessionID string, userID string) error {
	return s.store.Release(ctx, sessionID, userID)
}
