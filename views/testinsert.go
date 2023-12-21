package views

import (
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
