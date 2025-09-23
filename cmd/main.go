package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/viacheslav-korobeynikov/vacancy-site/config"
	"github.com/viacheslav-korobeynikov/vacancy-site/internal/home"
	"github.com/viacheslav-korobeynikov/vacancy-site/pkg/logger"
)

func main() {
	config.Init()                      // Получение данных из файла конфигурации
	config.NewDatabaseConfig()         // Вызов конфигурации БД
	logConfig := config.NewLogConfig() // Вызов конфигурации логов
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New() // Создание инстанса приложения Fiber

	app.Use(fiberzerolog.New()) // Middleware для логирования запросов

	app.Use(recover.New()) // Middleware, который перезапускает приложение в случае, если произошел вызов panic

	home.NewHandler(app, customLogger) // Добавили зависимость с хэндлером для главной страницы

	app.Listen(":3000") //Настраиваем порт, который будем слушать
}
