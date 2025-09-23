package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/viacheslav-korobeynikov/vacancy-site/pkg/templadapter"
	"github.com/viacheslav-korobeynikov/vacancy-site/views"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

// Функция конструктор
func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	// Группы роутов
	api := h.router.Group("/api") // Добавление группы роутов, вторым параметром можно добавить middleware ко всей группе
	api.Get("/", h.home)          // При Get запросе по адресу / вызываем функцию home
	api.Get("/error", h.error)
}

// Хэндлер для главной страницы
func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Hello("Anton")
	return templadapter.Render(c, component)
}

// Хэндлер для страницы error
func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 13).
		Msg("Информационный лог")
	return c.SendString("Error")
}
