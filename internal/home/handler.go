package home

import (
	"bytes"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
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
	tmpl, err := template.New("test").Parse("{{.Count}} - число пользователей") // Шаблон
	data := struct{ Count int }{Count: 1}                                       // Набор данных для подстановки
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Template error")
	}
	var tpl bytes.Buffer
	// Формируем шаблон
	if err := tmpl.Execute(&tpl, data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Template compile error")
	}
	return c.Send(tpl.Bytes())
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
