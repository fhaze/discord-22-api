package config

import (
	"os"
)

type Config struct {
	DbHost     string
	DbName     string
	DbUser     string
	DbPass     string
	RootApiKey string
	CommitHash string
}

var cfg *Config

func Instance() *Config {
	if cfg == nil {
		cfg = &Config{
			DbHost:     getEnvDefault("DB_HOST", "localhost:27017"),
			DbName:     getEnvDefault("DB_NAME", "discord-22"),
			DbUser:     getEnvDefault("DB_USER", "root"),
			DbPass:     getEnvDefault("DB_PASS", "password"),
			RootApiKey: getEnvDefault("ROOT_API_KEY", "22-key"),
			CommitHash: getEnvDefault("COMMIT_HASH", "Unknown"),
		}
	}
	return cfg
}

func getEnvDefault(env string, def string) (val string) {
	if val = os.Getenv(env); val == "" {
		val = def
	}
	return
}
