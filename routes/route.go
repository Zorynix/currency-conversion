package routes

import (
	"currency-conversion/repo"
	"currency-conversion/services"
	"currency-conversion/views"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
		logrus.Info("GET /v1/rates called")
		view := views.NewView(c, r.ratesService)
		err := view.ExchangeRateView()
		if err != nil {
			logrus.Errorf("Error in ExchangeRateView: %v", err)
			return err
		}
		logrus.Info("GET /v1/rates successful")
		return nil
	})

	v1.Get("/currencies", func(c *fiber.Ctx) error {
		logrus.Info("GET /v1/currencies called")
		view := views.NewView(c, r.ratesService)
		err := view.CurrenciesView()
		if err != nil {
			logrus.Errorf("Error in CurrenciesView: %v", err)
			return err
		}
		logrus.Info("GET /v1/currencies successful")
		return nil
	})

	v1.Get("/update", func(c *fiber.Ctx) error {
		logrus.Info("GET /v1/update called")
		view := views.NewView(c, r.ratesService)
		err := view.RateHistoryView()
		if err != nil {
			logrus.Errorf("Error in RateHistoryView: %v", err)
			return err
		}
		logrus.Info("GET /v1/update successful")
		return nil
	})
}
