package views

import (
	"currency-conversion/services"

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

	data, err := services.TestApi()
	if err != nil {
		log.Info().Err(err).Msg("")
		return fiber.NewError(fiber.StatusBadGateway)
	}

	return view.Ctx.SendString(string(data))
}
