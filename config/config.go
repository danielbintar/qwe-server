package config

import (
	"sync"

	"github.com/danielbintar/qwe-server/config/town"

	"github.com/go-redis/redis"
)

var (
	once sync.Once
	config *Config
)

type Config struct {
	townConfig  *town.TownConfig
	redisClient *redis.Client
}

func Instance() *Config {
	once.Do(func() {
		config = &Config{
			townConfig: town.Instance(),
			redisClient: RedisInstance(),
		}
	})

	return config
}
