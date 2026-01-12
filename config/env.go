package config

import (
	"fmt"
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)

	if ok {
		return value
	}
	return fallback
}

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", ":8000"),
		DBUser:     getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASSWORD", "admin"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "dev-db"),
	}
}

var Envs = initConfig()
