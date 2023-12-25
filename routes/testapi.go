package routes

import (
	"currency-conversion/views"

	"github.com/gofiber/fiber/v2"
)

func (route *Route) TestApiRoute() {
	route.Group.Get("/api", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c}
		return view.TestApiView()
	})
}
