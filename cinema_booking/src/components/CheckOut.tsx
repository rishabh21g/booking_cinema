import type { Session } from "../types/types";
import { Timer } from "./Timer";

export const Checkout = ({ session, status, onConfirm, onRelease, onTimerExpire }: { session: Session | null, status: { message: string, type: string } | null, onConfirm: () => void, onRelease: () => void, onTimerExpire: () => void }) => {
  if (status) {
    return (
      <div className="checkout">
        <div className={`status-msg ${status.type}`}>{status.message}</div>
      </div>
    );
  }

  if (!session) {
    return <div className="checkout"><div className="empty-state">Select an available seat to begin.</div></div>;
  }

  return (
    <div className="checkout">
      <h3>Checkout</h3>
      <div className="checkout-info"><span>Seat:</span> {session.seat_id}</div>
      <div className="checkout-info"><span>Movie:</span> {session.movie_id}</div>
      <div className="checkout-info"><span>Session:</span> {session.session_id.slice(0, 8)}&hellip;</div>
      <Timer expiresAt={session.expires_at} onExpire={onTimerExpire} />
      <div className="checkout-buttons">
        <button className="btn btn--confirm" onClick={onConfirm}>Confirm</button>
        <button className="btn btn--release" onClick={onRelease}>Release</button>
      </div>
    </div>
  );
};