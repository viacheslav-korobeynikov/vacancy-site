package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Инициазиация загрузки файла окружения
func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded")
}

// Структура конфигурации БД
type DatabaseConfig struct {
	url string
}

// Извление конфига БД из переменной окружения
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: os.Getenv("DATABASE_URL"),
	}
}
