package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	JWTKey  string
	AutoMin int
}

func getEnvInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return i
}

func Load() *Config {
	return &Config{
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASS"),
		DBHost:  os.Getenv("DB_HOST"),
		DBPort:  os.Getenv("DB_PORT"),
		DBName:  os.Getenv("DB_NAME"),
		AutoMin: getEnvInt("AUTO_COMPLETE_MINUTES", 5),
	}
}
