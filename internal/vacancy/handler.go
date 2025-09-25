package vacancy

import (
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/viacheslav-korobeynikov/vacancy-site/pkg/templadapter"
	"github.com/viacheslav-korobeynikov/vacancy-site/pkg/validator"
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
	form := VacancyCreateForm{
		Email: c.FormValue("email"),
	}
	//Добавили валидатор
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
	)
	time.Sleep(time.Second * 2)
	var component templ.Component
	// Если возникла хотя бы одна ошибка
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return templadapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return templadapter.Render(c, component)
}
