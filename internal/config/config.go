package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	TokenTTL   int64
	ENV        string
	LOG_LEVEL  string
}

func LoadConfig() (*Config, error) {

	tokenTTLStr := os.Getenv("TokenTTL")
	tokenTTL, err := strconv.ParseInt(tokenTTLStr, 10, 64)
	if err != nil {
		return nil, err
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		TokenTTL:   tokenTTL,
		ENV:        os.Getenv("APP_ENV"),
		LOG_LEVEL:  os.Getenv("LOG_LEVEL"),
	}, nil

}
