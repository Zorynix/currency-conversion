package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (view *View) CurrenciesView() error {

	log.Info().Msg("CurrenciesView called")
	dataCurrencies, err := view.MSQ.AddCurrencies()
	if err != nil {
		log.Error().Err(err).Msg("AddCurrencies")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(*dataCurrencies)
}

func (view *View) ExchangeRateView() error {

	log.Info().Msg("ExchangeRateView called")
	dataLatestExchangeRates, err := view.MSQ.AddRates()
	if err != nil {
		log.Error().Err(err).Msg("Error in AddRates")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(*dataLatestExchangeRates)
}

func (view *View) RateHistoryView() error {

	log.Info().Msg("RateHistoryView called")
	dataRateHistory, err := view.MSQ.UpdateRates()
	if err != nil {
		log.Error().Err(err).Msg("Error in UpdateRates")
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return view.Ctx.JSON(*dataRateHistory)
}
