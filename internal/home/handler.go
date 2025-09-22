package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

// Функция конструктор
func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	// Роутинг
	h.router.Get("/", h.home) // При Get запросе по адресу / вызываем функцию home
}

// Хэндлер для главной страницы
func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
