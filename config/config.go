package config

import (
	"sync"

	"github.com/danielbintar/qwe-server/config/town"
)

var (
	once sync.Once
	config *Config
)

type Config struct {
	townConfig *town.TownConfig
}

func Instance() *Config {
	once.Do(func() {
		config = &Config{
			townConfig: town.Instance(),
		}
	})

	return config
}
