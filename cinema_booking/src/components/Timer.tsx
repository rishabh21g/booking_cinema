import { useEffect, useState } from "react";

export const Timer = ({ expiresAt, onExpire }: { expiresAt: string, onExpire: () => void }) => {
  const [remaining, setRemaining] = useState(0);

  useEffect(() => {
    const calculateRemaining = () => Math.max(0, Math.floor((new Date(expiresAt).getTime() - Date.now()) / 1000));
    setRemaining(calculateRemaining());

    const interval = setInterval(() => {
      const newRemaining = calculateRemaining();
      setRemaining(newRemaining);
      if (newRemaining <= 0) {
        clearInterval(interval);
        onExpire();
      }
    }, 1000);

    return () => clearInterval(interval);
  }, [expiresAt, onExpire]);

  const mins = String(Math.floor(remaining / 60)).padStart(2, '0');
  const secs = String(remaining % 60).padStart(2, '0');

  return <div className={`timer ${remaining < 60 ? 'urgent' : ''}`}>{mins}:{secs}</div>;
};