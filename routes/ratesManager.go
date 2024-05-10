package routes

import (
	"currency-conversion/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (route *Route) ApiExchangeRateRoute() {
	route.Group.Get("/rates", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, MSQ: route.MSQ}
		log.Info().Interface("ctx", c).Msg("CTX")
		return view.ExchangeRateView()
	})
}

func (route *Route) ApiCurrenciesRoute() {
	route.Group.Get("/currencies", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, MSQ: route.MSQ}
		log.Info().Interface("ctx", c).Msg("CTX")
		return view.CurrenciesView()
	})
}

func (route *Route) ApiUpdateRates() {
	route.Group.Get("/update", func(c *fiber.Ctx) error {
		view := views.View{Ctx: c, MSQ: route.MSQ}
		log.Info().Interface("ctx", c).Msg("CTX")
		return view.RateHistoryView()
	})
}
