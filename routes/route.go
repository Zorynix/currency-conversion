package routes

import (
	"currency-conversion/repo"
	"currency-conversion/services"
	"currency-conversion/views"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App          *fiber.App
	ratesService services.RatesService
}

func NewRouter(app *fiber.App, db services.Database) *Router {
	currencyRepo := repo.NewCurrencyRepo(db.(*services.Mysql).DB)
	exchangeRatesRepo := repo.NewExchangeRatesRepo(db.(*services.Mysql).DB)
	rateHistoriesRepo := repo.NewRateHistoriesRepo(db.(*services.Mysql).DB)

	ratesService := services.NewRatesService(db, currencyRepo, exchangeRatesRepo, rateHistoriesRepo)
	return &Router{App: app, ratesService: ratesService}
}

func (r *Router) SetupRoutes() {
	v1 := r.App.Group("/v1")

	v1.Get("/rates", func(c *fiber.Ctx) error {
		view := views.NewView(c, r.ratesService)
		return view.ExchangeRateView()
	})

	v1.Get("/currencies", func(c *fiber.Ctx) error {
		view := views.NewView(c, r.ratesService)
		return view.CurrenciesView()
	})

	v1.Get("/update", func(c *fiber.Ctx) error {
		view := views.NewView(c, r.ratesService)
		return view.RateHistoryView()
	})
}
