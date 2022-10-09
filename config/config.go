package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App
	HTTP
	Log
	Mongo
	Secret
}

type App struct {
	Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"http_port" env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
}

type Mongo struct {
	URI        string `env-required:"true" yaml:"mongo_uri" env:"MONGO_URI"`
	Collection string `env-required:"true" yaml:"mongo_collection" env:"MONGO_COLLECTION"`
	Name       string `env-required:"true" yaml:"mongo_db_name" env:"MONGO_DB_NAME"`
}

type Secret struct {
	SignedKey string `env-required:"true" yaml:"signed_key" env:"SIGNED_KEY"`
}

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	fmt.Printf("%s %s\n", cfg.HTTP.Port, cfg.Log.Level)

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
