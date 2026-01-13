package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", ":8000"),
		DBUser:                 getEnv("DB_USER", "admin"),
		DBPassword:             getEnv("DB_PASSWORD", "admin"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "dev-db"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "CHANGE_ME"),
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)

	if ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	value, ok := os.LookupEnv(key)

	if ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

var Envs = initConfig()
