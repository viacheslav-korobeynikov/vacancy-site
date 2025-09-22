package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/viacheslav-korobeynikov/vacancy-site/config"
	"github.com/viacheslav-korobeynikov/vacancy-site/internal/home"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig() // Вызов конфигурации БД
	log.Println(dbConf)
	app := fiber.New() // Создание инстанса приложения Fiber

	app.Use(recover.New()) // Middleware, который перезапускает приложение в случае, если произошел вызов panic

	home.NewHandler(app) // Добавили зависимость с хэндлером для главной страницы

	app.Listen(":3000") //Настраиваем порт, который будем слушать
}
