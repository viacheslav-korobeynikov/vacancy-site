package config

import (
	"log"
	"os"
	"strconv"

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

// Работа с дефолтными значениями
func getString(key, defaultValue string) string {
	// Читаем из переменной окружения
	value := os.Getenv(key)
	// Если значение пустое
	if value == "" {
		// Возвращаем дефолтное значение
		return defaultValue
	}
	// Если значение не пустое - возвращаем значение переменной окружения
	return value
}

// Работа с числовыми значениями
func getInt(key string, defaultValue int) int {
	// Читаем из переменной окружения
	value := os.Getenv(key)
	// Конвертируем в int
	i, err := strconv.Atoi(value)
	// Если не удалось сконвертировать
	if err != nil {
		// Возвращаем дефолтное значение
		return defaultValue
	}
	// Если удалось сконвертировать - возвращаем значение int
	return i
}

// Работа с булево значениями
func getBool(key string, defaultValue bool) bool {
	// Читаем из переменной окружения
	value := os.Getenv(key)
	// Парсим значение bool
	b, err := strconv.ParseBool(value)
	// Если не удалось спарсить
	if err != nil {
		// Возвращаем дефолтное значение
		return defaultValue
	}
	// Если удалось сконвертировать - возвращаем значение bool
	return b
}

// Структура конфигурации БД
type DatabaseConfig struct {
	Url string
}

// Извление конфига БД из переменной окружения
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Url: getString("DATABASE_URL", ""),
	}
}

// Структура кофига логов
type LogConfig struct {
	Level  int
	Format string
}

// Извление конфига БД из переменной окружения
func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  getInt("LOG_LEVEL", 0),
		Format: getString("LOG_FORMAT", "json"),
	}
}
