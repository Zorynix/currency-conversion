package routes

import (
	"currency-conversion/views"

	"github.com/gofiber/fiber/v2"
)

func (route *Route) TestInsertRoute() {
	route.Group.Get("/insert", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, MSQ: route.MSQ}
		return view.TestInsertView()
	})
}
