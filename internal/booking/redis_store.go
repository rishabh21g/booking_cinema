package booking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const DEFALUT_HOLD_TTL = 2 * time.Minute

type RedisStore struct {
	rdb *redis.Client
}

// constructor func to create an instance of redis
func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{
		rdb: rdb,
	}
}

// reverse lookup for a session
func sessionkey(id string) string {
	return fmt.Sprintf("session:%s  ", id)
}

// booking a new seat
func (s *RedisStore) Book(b Booking) error {
	session, err := s.hold(b)

	if err != nil {
		return err
	}
	log.Printf("Session Booked: %v", session)
	return nil

}

// holding a seat for a time defined in ttl
func (s *RedisStore) hold(b Booking) (Booking, error) {
	id := uuid.New().String()
	now := time.Now()
	ctx := context.Background()
	key := fmt.Sprintf("seat:%s:%s", b.MoviedID, b.SeatID)
	b.ID = id

	val, _ := json.Marshal(b)

	res := s.rdb.SetArgs(ctx, key, val, redis.SetArgs{
		Mode: "NX",
		TTL:  DEFALUT_HOLD_TTL,
	})

	ok := res.Val() == "OK"
	if !ok {
		return Booking{}, ErrSeatsAlreadyBooked
	}

	s.rdb.Set(ctx, sessionkey(id), key, DEFALUT_HOLD_TTL)

	return Booking{
		MoviedID:  b.MoviedID,
		UserID:    b.UserID,
		Status:    "Hold",
		ID:        id,
		SeatID:    b.SeatID,
		ExpiresAt: now.Add(DEFALUT_HOLD_TTL),
	}, nil

}

func parseSession(val string) (Booking, error) {
	var data Booking
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return Booking{}, err
	}

	return Booking{
		MoviedID:  data.MoviedID,
		SeatID:    data.SeatID,
		UserID:    data.UserID,
		ID:        data.ID,
		Status:    data.Status,
		ExpiresAt: data.ExpiresAt,
	}, nil
}

// confirm the booking and remove the ttl with permanent
func (s *RedisStore) Confirm(ctx context.Context, sessionID string, userID string) (Booking, error) {
	session, sk, err := s.getSession(ctx, sessionID, userID)
	if err != nil {
		return Booking{}, err
	}
	s.rdb.Persist(ctx, sk)
	s.rdb.Persist(ctx, sessionkey(sessionID))
	session.Status = "Confimed"
	data := Booking{
		MoviedID: session.MoviedID,
		UserID:   session.UserID,
		Status:   "confirmed",
		ID:       session.ID,
		SeatID:   session.SeatID,
	}
	val, err := json.Marshal(data)
	if err != nil {
		return Booking{}, err
	}
	s.rdb.Set(ctx, sk, val, 0)
	return session, nil
}

// Get the session
func (s *RedisStore) getSession(ctx context.Context, sessionID string, userID string) (Booking, string, error) {

	sk, err := s.rdb.Get(ctx, sessionkey(sessionID)).Result()
	if err != nil {
		return Booking{}, "", err
	}
	val, err := s.rdb.Get(ctx, sk).Result()
	if err != nil {
		return Booking{}, "", err
	}
	session, err := parseSession(val)

	if err != nil {
		return Booking{}, "", err
	}
	return session, sk, nil
}

// release the seat
func (s *RedisStore) Release(ctx context.Context, sessioID string, userID string) error {
	_, sk, err := s.getSession(ctx, sessioID, userID)
	if err != nil {
		return err
	}
	s.rdb.Del(ctx, sk, sessionkey(sessioID))
	return nil
}

func (s *RedisStore) ListBookings(movieID string) []Booking {
	pattern := fmt.Sprintf("seat:%s:*", movieID)
	var sessions []Booking

	ctx := context.Background()

	iter := s.rdb.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		val, err := s.rdb.Get(ctx, iter.Val()).Result()
		if err != nil {
			continue
		}
		session, err := parseSession(val)
		if err != nil {
			continue
		}
		sessions = append(sessions, session)
	}

	return sessions
}
