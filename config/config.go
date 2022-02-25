package config

import (
	"os"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPass     string
	RootApiKey string
}

var cfg *Config

func Instance() *Config {
	if cfg == nil {
		cfg = &Config{
			DbHost:     getEnvDefault("DB_HOST", "localhost"),
			DbPort:     getEnvDefault("DB_PORT", "27017"),
			DbName:     getEnvDefault("DB_NAME", "eaglejump"),
			DbUser:     getEnvDefault("DB_USER", "root"),
			DbPass:     getEnvDefault("DB_PASS", "password"),
			RootApiKey: getEnvDefault("ROOT_API_KEY", "eagle-jump-key"),
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
