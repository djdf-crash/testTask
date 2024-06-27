package config

import (
	"github.com/caarlos0/env/v11"
	"time"
)

type ServerConfig struct {
	Port            string        `env:"PORT" envDefault:"9000"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"5s"`
}

type DatabaseConfig struct {
	DSN            string `env:"DSN,required"`
	NeedRecreateDB bool   `env:"NEED_RECREATE" envDefault:"true"`
}

type Config struct {
	ServerConfig   ServerConfig   `envPrefix:"SERVER_"`
	DatabaseConfig DatabaseConfig `envPrefix:"DATABASE_"`
}

func ReadConfig() (Config, error) {
	var (
		cfg Config
		err error
	)
	err = env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
