package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		// Addr:     "localhost:6379", //redis地址
		// Password: "",               //redis密碼
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   0, //默認資料庫，預設為0
	})
	pong, err := client.Ping(Ctx).Result()
	log.Println("Redis ping:", pong)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
