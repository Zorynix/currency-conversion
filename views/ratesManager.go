package views

import (
	"currency-conversion/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (view *View) TestInsertView() error {

	err := view.MSQ.TestInsert()
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadGateway)
	}
	return err
}

func (view *View) TestApiView() error {

	dataAllCurrencies, err := view.MSQ.TestApiAllCurrencies()
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	dataLatestExchangeRates, err := view.MSQ.LatestExchangeRates()
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.JSON(dto.AllData{
		DataAllCurrencies:       *dataAllCurrencies,
		DataLatestExchangeRates: *dataLatestExchangeRates,
	})
}
