package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"p3ld3v.dev/template/app/domain"
)

func LoadConfig() domain.Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARN: .env file not found")
	}

	env := domain.Config{
		Port:         getEnvKeyStr("PORT", "8080"),
		Host:         getEnvKeyStr("HOST", "0.0.0.0"),
		LogLevel:     getEnvKeyStr("LOG_LEVEL", "debug"),
		DbConnection: getEnvKeyStr("DB_CONNECTION_STRING", ""),
	}
	return env
}

func getEnvKeyStr(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
