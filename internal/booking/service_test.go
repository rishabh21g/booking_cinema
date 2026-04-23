package booking

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/google/uuid"
)

func TestConcurrentBokking(t *testing.T) {
	// creating a in memory store
	store := NewMemoryStrore()

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
			err := service_store.Book(Booking{
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

//   go test -race ./... -v
//          github.com/rishabh21g/booking_cinema/cmd        [no test files]
// === RUN   TestConcurrentBokking
// ==================
// WARNING: DATA RACE
// Read at 0x00c0000a0020 by goroutine 15:
//   github.com/rishabh21g/booking_cinema/internal/booking.(*MemoryStore).Book()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/memory.go:17 +0x64
//   github.com/rishabh21g/booking_cinema/internal/booking.(*Service).Book()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service.go:12 +0x174
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking.func1()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:36 +0xcc
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking.gowrap1()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:47 +0x38

// Previous write at 0x00c0000a0020 by goroutine 12:
//   github.com/rishabh21g/booking_cinema/internal/booking.(*MemoryStore).Book()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/memory.go:21 +0x9c
//   github.com/rishabh21g/booking_cinema/internal/booking.(*Service).Book()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service.go:12 +0x174
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking.func1()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:36 +0xcc
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking.gowrap1()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:47 +0x38

// Goroutine 15 (running) created at:
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:32 +0x150
//   testing.tRunner()
//       /usr/local/go/src/testing/testing.go:2036 +0x164
//   testing.(*T).Run.gowrap1()
//       /usr/local/go/src/testing/testing.go:2101 +0x34

// Goroutine 12 (finished) created at:
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:32 +0x150
//   testing.tRunner()
//       /usr/local/go/src/testing/testing.go:2036 +0x164
//   testing.(*T).Run.gowrap1()
//       /usr/local/go/src/testing/testing.go:2101 +0x34
// ==================
// ==================
// WARNING: DATA RACE
// Read at 0x00c0001a4870 by goroutine 9:
//   runtime.mapaccess1_faststr()
//       /usr/local/go/src/internal/runtime/maps/runtime_faststr.go:101 +0x28c
//   github.com/rishabh21g/booking_cinema/internal/booking.(*MemoryStore).Book()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/memory.go:17 +0x58
//   github.com/rishabh21g/booking_cinema/internal/booking.(*Service).Book()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service.go:12 +0x174
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking.func1()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:36 +0xcc
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking.gowrap1()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:47 +0x38

// Previous write at 0x00c0001a4870 by goroutine 12:
//   runtime.mapaccess2_faststr()
//       /usr/local/go/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
//   github.com/rishabh21g/booking_cinema/internal/booking.(*MemoryStore).Book()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/memory.go:21 +0x8c
//   github.com/rishabh21g/booking_cinema/internal/booking.(*Service).Book()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service.go:12 +0x174
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking.func1()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:36 +0xcc
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking.gowrap1()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:47 +0x38

// Goroutine 9 (running) created at:
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:32 +0x150
//   testing.tRunner()
//       /usr/local/go/src/testing/testing.go:2036 +0x164
//   testing.(*T).Run.gowrap1()
//       /usr/local/go/src/testing/testing.go:2101 +0x34

// Goroutine 12 (finished) created at:
//   github.com/rishabh21g/booking_cinema/internal/booking.TestConcurrentBokking()
//       /Users/rishabh23g/Development/booking_cinema/internal/booking/service_test.go:32 +0x150
//   testing.tRunner()
//       /usr/local/go/src/testing/testing.go:2036 +0x164
//   testing.(*T).Run.gowrap1()
//       /usr/local/go/src/testing/testing.go:2101 +0x34
// ==================
//     service_test.go:53: SUCCESSES: 1 , FAILURES 99999
//     testing.go:1712: race detected during execution of test
// --- FAIL: TestConcurrentBokking (0.35s)
// FAIL
// FAIL    github.com/rishabh21g/booking_cinema/internal/booking   1.110s
// FAIL
