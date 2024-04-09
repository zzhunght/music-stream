package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(url string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:             url,
		Password:         "", // no password set
		DB:               0,  // use default DB
		DisableIndentity: true,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	fmt.Println("Connected to Redis:", pong)
	return rdb
}
