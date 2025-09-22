package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // Создание инстанса приложения Fiber

	//Простейший хэндлер
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello") // Возвращаем ответ
	})

	app.Listen(":3000") //Настраиваем порт, который будем слушать
}
