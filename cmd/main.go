package main

import (
	"log"
	"net/http"

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

	// if error just log it
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("GET /movies", listMovies)

}
