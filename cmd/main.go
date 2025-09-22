package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viacheslav-korobeynikov/vacancy-site/internal/home"
)

func main() {
	app := fiber.New() // Создание инстанса приложения Fiber

	home.NewHandler(app) // Добавили зависимость с хэндлером для главной страницы

	app.Listen(":3000") //Настраиваем порт, который будем слушать
}
