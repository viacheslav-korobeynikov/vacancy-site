package vacancy

import (
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
		Email:    c.FormValue("email"),
		Role:     c.FormValue("role"),
		Type:     c.FormValue("type"),
		Company:  c.FormValue("company"),
		Location: c.FormValue("location"),
		Salary:   c.FormValue("salary"),
	}
	//Добавили валидатор
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "Расположение не задано"},
		&validators.StringIsPresent{Name: "Role", Field: form.Role, Message: "Не указана должность"},
		&validators.StringIsPresent{Name: "Company", Field: form.Company, Message: "Не указано название компании"},
		&validators.StringIsPresent{Name: "Type", Field: form.Type, Message: "Не указана сфера деятельности компании"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "Не указана зарплата"},
	)
	var component templ.Component
	// Если возникла хотя бы одна ошибка
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return templadapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return templadapter.Render(c, component)
}
