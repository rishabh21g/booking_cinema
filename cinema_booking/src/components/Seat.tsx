import { useMemo } from "react";
import type { Movie, SeatStatus } from "../types/types";
import Legend from "./Legend";

const SeatGrid = ({ movie, seatStatuses, userID, onHoldSeat }: { movie: Movie, seatStatuses: Record<string, SeatStatus>, userID: string, onHoldSeat: (seatId: string) => void }) => {
  const rows = useMemo(() => {
    const rowLabels = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    return Array.from({ length: movie.rows }, (_, r) => {
      const seats = Array.from({ length: movie.seats_per_row }, (_, s) => {
        const seatID = rowLabels[r] + (s + 1);
        const status = seatStatuses[seatID];
        let className = 'seat';
        if (status) {
          if (status.confirmed) className += ' seat--confirmed';
          else if (status.booked && status.user_id === userID) className += ' seat--held-mine';
          else if (status.booked) className += ' seat--held-other';
        }
        return { id: seatID, number: s + 1, className };
      });
      return { label: rowLabels[r], seats };
    });
  }, [movie, seatStatuses, userID]);

  return (
    <div className="screen-area">
      <div className="screen-label">Screen</div>
      <div className="screen-bar"></div>
      <div className="seat-grid">
        {rows.map(row => (
          <div key={row.label} className="seat-row">
            <div className="row-label">{row.label}</div>
            {row.seats.map(seat => (
              <button key={seat.id} className={seat.className} onClick={() => onHoldSeat(seat.id)}>
                {seat.number}
              </button>
            ))}
            <div className="row-label">{row.label}</div>
          </div>
        ))}
      </div>
      <Legend />
    </div>
  );
};

export default SeatGrid