package vacancy

import (
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
	h.customLogger.Info().Msg(email)
	component := components.Notification("Вакансия успешно создана")
	return templadapter.Render(c, component)
}
