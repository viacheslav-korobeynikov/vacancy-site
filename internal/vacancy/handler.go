package vacancy

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/viacheslav-korobeynikov/vacancy-site/pkg/templadapter"
	"github.com/viacheslav-korobeynikov/vacancy-site/views/components"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy") // Создаем группу роутов
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")
	var component templ.Component
	if email == "" {
		component = components.Notification("Не указан email", components.NotificationFail)
		return templadapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return templadapter.Render(c, component)
}
