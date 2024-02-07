package views

import (
	"currency-conversion/services"

	"github.com/gofiber/fiber/v2"
)

type View struct {
	Ctx *fiber.Ctx
	MSQ *services.Mysql
	App *fiber.App
}
