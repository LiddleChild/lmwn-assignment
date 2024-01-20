package config

import (
	"os"
)

type Config struct {
	appMode   string
	AppPort   string
	CovidHost string
}

func (c *Config) IsDevelopment() bool {
	return c.appMode == "development"
}

func Load() *Config {
	config := new(Config)

	config.appMode = env("APP_ENV", "production")
	config.AppPort = env("APP_PORT", "8000")
	config.CovidHost = env("COVID_STAT_SERVER", "")

	return config
}

func env(key string, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}

	return val
}
