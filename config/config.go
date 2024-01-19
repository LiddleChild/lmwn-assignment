package config

import (
	"os"
)

type Config struct {
	mode    string
	AppPort string
}

func (c *Config) IsDevelopment() bool {
	return c.mode == "development"
}

func Load() *Config {
	config := new(Config)

	config.mode = env("APP_ENV", "production")
	config.AppPort = env("APP_PORT", "8000")

	return config
}

func env(key string, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}

	return val
}
