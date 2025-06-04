package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBDriver string
	DBSource string
	Port     string
}

func Load() *Config {
	return &Config{
		DBDriver: getEnv("DB_DRIVER", "sqlite3"),
		DBSource: getEnv("DB_SOURCE", "app.db"),
		Port:     getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func (c *Config) Addr() string {
	return fmt.Sprintf(":%s", c.Port)
}
