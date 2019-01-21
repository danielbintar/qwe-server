package config

import (
	"os"
	"sync"
	"github.com/go-redis/redis"
)

var (
	once2 sync.Once
	redisClient *redis.Client
)

func RedisInstance() *redis.Client {
	once2.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		})

		_, err := redisClient.Ping().Result()
		if err != nil { panic(err) }
	})

	return redisClient
}
