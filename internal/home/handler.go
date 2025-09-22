package home

import (
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
}

// /api/
// /api/error

// Функция конструктор
func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	// Группы роутов
	api := h.router.Group("/api") // Добавление группы роутов, вторым параметром можно добавить middleware ко всей группе
	api.Get("/", h.home)          // При Get запросе по адресу / вызываем функцию home
	api.Get("/error", h.error)
}

// Хэндлер для главной страницы
func (h *HomeHandler) home(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}

// Хэндлер для страницы error
func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
