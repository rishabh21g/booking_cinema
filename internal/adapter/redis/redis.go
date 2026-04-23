package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(address string) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: address,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("REDIS PING % v", err)
	}
	log.Printf("connected to redis: %s", address)
	return rdb
}
