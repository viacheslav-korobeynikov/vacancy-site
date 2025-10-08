package templadapter

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// Адаптер между Templ и Fiber
func Render(c *fiber.Ctx, component templ.Component, code int) error {
	return adaptor.HTTPHandler(templ.Handler(component, templ.WithStatus(code)))(c)
}
