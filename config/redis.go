package config

import (
	"os"
	"github.com/go-redis/redis"
)

func RedisInstance() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	_, err := client.Ping().Result()
	if err != nil { panic(err) }

	return client
}
