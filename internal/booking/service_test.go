package booking

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/google/uuid"
	"github.com/rishabh21g/booking_cinema/internal/adapter/redis"
)

func TestConcurrentBokking(t *testing.T) {
	// creating a in memory store
	// store := NewConcurrentMemoryStore() // lock safe memory
	// store := NewMemoryStrore()  not safe memory
	store := NewRedisStore(redis.NewRedisClient("localhost:6379"))
	// creating a service
	service_store := NewService(store)

	//no of concurrent threads going to access the critical section mean seatID
	const NUMBER_OF_GOROUTINES = 100000

	// defining the atomicity of processes
	var (
		Successes atomic.Int64
		Failures  atomic.Int64
		wg        sync.WaitGroup
	)

	// adding the goroutines in the waitgroup
	wg.Add(NUMBER_OF_GOROUTINES)

	for i := range NUMBER_OF_GOROUTINES {
		go func(ID int) {
			defer wg.Done()

			// booking on described seat
			_, err := service_store.Book(Booking{
				MoviedID: "screen-1",
				SeatID:   "C2",
				UserID:   uuid.New().String(),
			})

			if err == nil {
				Successes.Add(1)
			} else {
				Failures.Add(1)
			}
		}(i)
	}

	// wait for all goroutines to finishes their jobs
	wg.Wait()

	t.Logf("SUCCESSES: %d , FAILURES %d", Successes.Load(), Failures.Load())
	if Successes.Load() != 1 {
		t.Errorf("expected 1 success, but got %d", Successes.Load())
	}
	if Failures.Load() != NUMBER_OF_GOROUTINES-1 {
		t.Errorf("expected %d failed booking, but got %d", NUMBER_OF_GOROUTINES-1, Failures.Load())
	}
}
