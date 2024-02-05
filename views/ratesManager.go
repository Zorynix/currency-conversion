package views

import (
	"currency-conversion/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (view *View) ApiView() error {

	dataAllCurrencies, err := view.MSQ.InsertAllCurrencies()
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	dataLatestExchangeRates, err := view.MSQ.InsertLatestExchangeRates()
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(dto.AllData{
		DataAllCurrencies:       *dataAllCurrencies,
		DataLatestExchangeRates: *dataLatestExchangeRates,
	})
}
