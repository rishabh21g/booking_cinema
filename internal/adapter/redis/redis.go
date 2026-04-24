package redis

import (
	"context"
	"log"

	goredis "github.com/redis/go-redis/v9"
)

func NewRedisClient(address string) *goredis.Client {

	rdb := goredis.NewClient(&goredis.Options{
		Addr: address,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("REDIS PING % v", err)
	}
	log.Printf("connected to redis: %s", address)
	return rdb
}
