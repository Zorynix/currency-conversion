package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (view *View) CurrenciesView() error {
	logrus.Info("CurrenciesView called")
	data, err := view.ratesService.AddCurrencies()
	if err != nil {
		logrus.Errorf("Failed to add currencies: %v", err)
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return view.Ctx.JSON(*data)
}

func (view *View) ExchangeRateView() error {
	logrus.Info("ExchangeRateView called")
	data, err := view.ratesService.AddRates()
	if err != nil {
		logrus.Errorf("Error in AddRates: %v", err)
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return view.Ctx.JSON(*data)
}

func (view *View) RateHistoryView() error {
	logrus.Info("RateHistoryView called")
	message, err := view.ratesService.UpdateRates()
	if err != nil {
		logrus.Errorf("Error in UpdateRates: %v", err)
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return view.Ctx.SendString(message)
}
