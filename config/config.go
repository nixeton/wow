package config

import (
	cl "github.com/nixeton/ca"
	"time"
)

type (
	// Config -.
	Config struct {
		TCP `yaml:"tpc"`
		Log `yaml:"logger"`
		Pow `yaml:"pow"`
	}

	TCP struct {
		Address   string        `env-required:"true" env:"TCP_ADDRESS"`
		KeepAlive time.Duration `env-required:"true" env:"TCP_KEEP_ALIVE"`
		Deadline  time.Duration `env-required:"true" env:"TCP_DEADLINE"`
	}

	Pow struct {
		Difficulty int `env-required:"true" env:"POW_DIFFICULTY"`
	}

	Log struct {
		Level string `env-required:"true"   env:"LOG_LEVEL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cl.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
