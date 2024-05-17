package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (view *View) CurrenciesView() error {
	log.Info().Msg("CurrenciesView called")
	data, err := view.ratesService.AddCurrencies()
	if err != nil {
		log.Error().Err(err).Msg("AddCurrencies")
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return view.Ctx.JSON(*data)
}

func (view *View) ExchangeRateView() error {
	log.Info().Msg("ExchangeRateView called")
	data, err := view.ratesService.AddRates()
	if err != nil {
		log.Error().Err(err).Msg("Error in AddRates")
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return view.Ctx.JSON(*data)
}

func (view *View) RateHistoryView() error {
	log.Info().Msg("RateHistoryView called")
	message, err := view.ratesService.UpdateRates()
	if err != nil {
		log.Error().Err(err).Msg("Error in UpdateRates")
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return view.Ctx.SendString(message)
}
