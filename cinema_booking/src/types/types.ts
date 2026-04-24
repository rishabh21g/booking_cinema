// --- Type Definitions ---
export interface Movie {
  id: string;
  title: string;
  rows: number;
  seats_per_row: number;
}

export interface SeatStatus {
  seat_id: string;
  booked: boolean;
  confirmed: boolean;
  user_id?: string;
}

export interface Session {
  session_id: string;
  movie_id: string;
  seat_id: string;
  expires_at: string;
}