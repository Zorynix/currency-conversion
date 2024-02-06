package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (view *View) CurrenciesView() error {

	dataCurrencies, err := view.MSQ.InsertCurrencies()
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(*dataCurrencies)
}

func (view *View) ExchangeRateView() error {

	dataLatestExchangeRates, err := view.MSQ.InsertExchangeRates()
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(*dataLatestExchangeRates)
}

func (view *View) RateHistoryView() error {

	dataRateHistory, err := view.MSQ.UpdateRates()
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return view.Ctx.JSON(*dataRateHistory)
}
