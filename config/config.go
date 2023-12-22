package config

import (
	"fmt"
	cl "github.com/nixeton/ca"
	"time"
)

type (
	// Config -.
	Config struct {
		App `yaml:"app"`
		TCP `yaml:"tpc"`
		Log `yaml:"logger"`
		Pow `yaml:"pow"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	TCP struct {
		Address   string        `env-required:"true" yaml:"address" env:"TCP_ADDRESS"`
		KeepAlive time.Duration `env-required:"true" yaml:"keep_alive" env:"TCP_KEEP_ALIVE"`
		Deadline  time.Duration `env-required:"true" yaml:"deadline" env:"TCP_DEADLINE"`
	}

	Pow struct {
		Difficulty int `env-required:"true" yaml:"difficulty" env:"POW_DIFFICULTY"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cl.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cl.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
