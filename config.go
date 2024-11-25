package main

import "os"

type Config struct {
	BaseURL string
}

func NewConfigFromEnv() *Config {
	return &Config{
		BaseURL: envOr("PREDBOT_BASE_URL", "https://omeda.city"),
	}
}

func envOr(key, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		val = def
	}

	return val
}
