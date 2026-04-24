import { useState, useEffect, useCallback } from 'react';
import './App.css';
import type { Movie, SeatStatus, Session } from './types/types';
import { Header } from './components/Header';
import MovieList from './components/Movies';
import { Checkout } from './components/CheckOut';
import SeatGrid from './components/Seat';

// --- API Helper ---
async function api<T>(method: string, path: string, body?: unknown): Promise<T | null> {
  const opts: RequestInit = {
    method,
    headers: { 'Content-Type': 'application/json' },
  };
  if (body) {
    opts.body = JSON.stringify(body);
  }
  const response = await fetch(path, opts);

  if (response.status === 204) {
    return null;
  }

  const data = await response.json();
  if (!response.ok) {
    throw new Error(data.error || 'Request failed');
  }
  return data;
}

function App() {
  const [userID] = useState(() => crypto.randomUUID().replace(/-/g, '').slice(0, 12));
  const [movies, setMovies] = useState<Movie[]>([]);
  const [selectedMovie, setSelectedMovie] = useState<Movie | null>(null);
  const [seatStatuses, setSeatStatuses] = useState<Record<string, SeatStatus>>({});
  const [activeSession, setActiveSession] = useState<Session | null>(null);
  const [checkoutStatus, setCheckoutStatus] = useState<{ message: string; type: 'success' | 'error' } | null>(null);

  // --- Effects ---
  useEffect(() => {
    api<Movie[]>('GET', '/movies').then(data => setMovies(data || []));
  }, []);

  const fetchSeats = useCallback(() => {
    if (!selectedMovie) return;
    api<SeatStatus[]>('GET', `/movies/${selectedMovie.id}/seats`).then(seats => {
      const statusMap = (seats || []).reduce((acc, s) => {
        acc[s.seat_id] = s;
        return acc;
      }, {} as Record<string, SeatStatus>);
      setSeatStatuses(statusMap);
    });
  }, [selectedMovie]);

  useEffect(() => {
    if (!selectedMovie) return;
    fetchSeats();
    const pollInterval = setInterval(fetchSeats, 2000);
    return () => clearInterval(pollInterval);
  }, [selectedMovie, fetchSeats]);

  // --- Handlers ---
  const handleSelectMovie = (movie: Movie) => {
    if (activeSession) {
      api('DELETE', `/sessions/${activeSession.session_id}`, { user_id: userID }).catch(() => {});
    }
    setActiveSession(null);
    setSelectedMovie(movie);
  };

  const handleHoldSeat = async (seatID: string) => {
    if (activeSession) return;
    try {
      const session = await api<Session>('POST', `/movies/${selectedMovie!.id}/seats/${seatID}/hold`, { user_id: userID });
      if (session) {
        setActiveSession(session);
        fetchSeats();
      }
    } catch (err) {
      showTempCheckoutStatus((err as Error).message, 'error');
    }
  };

  const handleConfirmSeat = async () => {
    if (!activeSession) return;
    try {
      await api('PUT', `/sessions/${activeSession.session_id}/confirm`, { user_id: userID });
      setActiveSession(null);
      fetchSeats();
      showTempCheckoutStatus('Confirmed!', 'success');
    } catch (err) {
      showTempCheckoutStatus((err as Error).message, 'error');
    }
  };

  const handleReleaseSeat = async () => {
    if (!activeSession) return;
    try {
      await api('DELETE', `/sessions/${activeSession.session_id}`, { user_id: userID });
      setActiveSession(null);
      fetchSeats();
    } catch (err) {
      showTempCheckoutStatus((err as Error).message, 'error');
    }
  };

  const showTempCheckoutStatus = (message: string, type: 'success' | 'error') => {
    setCheckoutStatus({ message, type });
    setTimeout(() => {
      setCheckoutStatus(null);
    }, 3000);
  };

  const onTimerExpire = () => {
    setActiveSession(null);
    fetchSeats();
    showTempCheckoutStatus('Hold expired', 'error');
  };

  return (
    <>
      <Header userID={userID} />
      <MovieList movies={movies} selectedMovie={selectedMovie} onSelectMovie={handleSelectMovie} />
      {selectedMovie && (
        <div className="content">
          <SeatGrid
            movie={selectedMovie}
            seatStatuses={seatStatuses}
            userID={userID}
            onHoldSeat={handleHoldSeat}
          />
          <Checkout
            session={activeSession}
            status={checkoutStatus}
            onConfirm={handleConfirmSeat}
            onRelease={handleReleaseSeat}
            onTimerExpire={onTimerExpire}
          />
        </div>
      )}
    </>
  );
}

export default App;