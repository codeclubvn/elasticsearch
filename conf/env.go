package conf

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	Port     string `env:"PORT" envDefault:"8082"`
	ESDomain string `env:"ES_DOMAIN" envDefault:"http://localhost:9200"`
}

var cfg Config

func SetEnv() {
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("Failed to read environment variables: %v", err)
		return
	}
}

func GetEnv() Config {
	return cfg
}
