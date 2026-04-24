import type { Movie } from "../types/types";

const MovieList = ({ movies, selectedMovie, onSelectMovie }: { movies: Movie[], selectedMovie: Movie | null, onSelectMovie: (movie: Movie) => void }) => (
  <div className="movies">
    {movies.map(m => (
      <div
        key={m.id}
        className={`movie-card ${selectedMovie?.id === m.id ? 'selected' : ''}`}
        onClick={() => onSelectMovie(m)}
      >
        <h3>{m.title}</h3>
        <p>{m.rows} rows &times; {m.seats_per_row} seats</p>
      </div>
    ))}
  </div>
);

export default MovieList


