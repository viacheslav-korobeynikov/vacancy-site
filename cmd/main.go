package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/viacheslav-korobeynikov/vacancy-site/config"
	"github.com/viacheslav-korobeynikov/vacancy-site/internal/home"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig() // Вызов конфигурации БД
	log.Println(dbConf)
	app := fiber.New() // Создание инстанса приложения Fiber

	home.NewHandler(app) // Добавили зависимость с хэндлером для главной страницы

	app.Listen(":3000") //Настраиваем порт, который будем слушать
}
