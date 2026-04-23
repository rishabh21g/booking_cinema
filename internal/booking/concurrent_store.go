package booking

import "sync"

type ConcurrentMemoryStore struct {
	bookings map[string]Booking // B2 Seat ID (string)
	sync.RWMutex
}

// constructor function rerturning the pointer
func NewConcurrentMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (m *ConcurrentMemoryStore) Book(b Booking) error {

	// thread safe critical section using mutex
	m.Lock()
	defer m.Unlock() // unclock before ending the func

	// check the seat is booked or not
	if _, exits := m.bookings[b.SeatID]; exits {
		return ErrSeatsAlreadyBooked
	}
	// seat booked
	m.bookings[b.SeatID] = b
	return nil
}

func (m *ConcurrentMemoryStore) ListBookings(movieID string) []Booking {
	m.RLock()         // read lock for concurrent reading
	defer m.RUnlock() // then unlock

	// creating an array to return
	var results []Booking

	// returning all the list of bookings by looping over the memory
	for _, b := range m.bookings {
		if b.MoviedID == movieID {
			results = append(results, b)
		}
	}

	return results

}
