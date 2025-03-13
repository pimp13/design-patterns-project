package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	PublicHost            string
	Port                  string
	DBUser                string
	DBPassword            string
	DBAddress             string
	DBName                string
	JWTExpirationInSecond int64
	JWTKey                string
}

var Envs = initConfig()

func initConfig() *Config {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		PublicHost:            getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                  getEnv("APP_PORT", ":8080"),
		DBUser:                getEnv("DB_USER", "root"),
		DBPassword:            getEnv("DB_PASSWORD", "root"),
		DBAddress:             fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                getEnv("DB_NAME", "db_name"),
		JWTExpirationInSecond: getEnvAsInt("JWT_EXPIRATION_IN_SECOND", 3600*24*7),
		JWTKey:                getEnv("JWT_KEY", "not-secret-key-create-secret-key"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
