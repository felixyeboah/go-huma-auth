package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"huma-auth/config"
	"log"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	// env config
	env, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal(err)
	}

	options, err := redis.ParseURL(env.RedisURL)
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(options)

	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")

	return client
}
