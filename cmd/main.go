package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rishabh21g/booking_cinema/internal/adapter/redis"
	"github.com/rishabh21g/booking_cinema/internal/booking"
	"github.com/rishabh21g/booking_cinema/utils"
)

type movieResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}

var movies = []movieResponse{
	{ID: "khatta-meetha", Title: "Khatta Meetha", Rows: 5, SeatsPerRow: 8},
	{ID: "Bhagam Bhaag", Title: "Bhagam Bhag", Rows: 4, SeatsPerRow: 6},
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, movies)
}

func main() {
	// creating a new server using multiplexer
	mux := http.NewServeMux()

	// new redis store
	store := booking.NewRedisStore(redis.NewRedisClient("localhost:6379"))
	service_store := booking.NewService(store)
	bookingHandler := booking.NewHandler(service_store)

	// Register all your handlers first
	mux.HandleFunc("GET /movies", listMovies)
	mux.HandleFunc("GET /movies/{movieID}/seats", bookingHandler.ListSeats)
	mux.HandleFunc("POST /movies/{movieID}/seats/{seatID}/hold", bookingHandler.HoldSeat)
	mux.HandleFunc("PUT /sessions/{sessionID}/confirm", bookingHandler.ConfirmSession)
	mux.HandleFunc("DELETE /sessions/{sessionID}", bookingHandler.ReleaseSession)

	fmt.Println("Server is listening on port : 3000")

	// Start the server last, as it's a blocking call
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
