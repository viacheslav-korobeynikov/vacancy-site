package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/viacheslav-korobeynikov/vacancy-site/config"
	"github.com/viacheslav-korobeynikov/vacancy-site/internal/home"
)

func main() {
	config.Init()                      // Получение данных из файла конфигурации
	config.NewDatabaseConfig()         // Вызов конфигурации БД
	logConfig := config.NewLogConfig() // Вызов конфигурации логов

	app := fiber.New() // Создание инстанса приложения Fiber

	log.SetLevel(log.Level(logConfig.Level)) // Устанавливаем уровень логов в зависимости от окружения

	app.Use(logger.New()) // Middleware для логирования запросов

	app.Use(recover.New()) // Middleware, который перезапускает приложение в случае, если произошел вызов panic

	home.NewHandler(app) // Добавили зависимость с хэндлером для главной страницы

	app.Listen(":3000") //Настраиваем порт, который будем слушать
}
