package booking

type MemoryStore struct {
	bookings map[string]Booking // B2 Seat ID (string)
}

// constructor function rerturning the pointer
func NewMemoryStrore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (m *MemoryStore) Book(b Booking) error {

	// check the seat is booked or not
	if _, exits := m.bookings[b.SeatID]; exits {
		return ErrSeatsAlreadyBooked
	}
	// seat booked
	m.bookings[b.SeatID] = b
	return nil
}

func (m *MemoryStore) ListBookings(movieID string) []Booking {
	var results []Booking

	// returning all the list of bookings by looping over the memory
	for _, b := range m.bookings {
		if b.MoviedID == movieID {
			results = append(results, b)
		}
	}

	return results

}
