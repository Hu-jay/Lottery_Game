package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	pong, err := client.Ping(Ctx).Result()
	log.Println("Redis ping:", pong)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
