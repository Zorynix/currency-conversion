package views

import (
	"currency-conversion/internal/services"

	"github.com/gofiber/fiber/v2"
)

type View struct {
	Ctx          *fiber.Ctx
	ratesService services.RatesService
}

func NewView(ctx *fiber.Ctx, ratesService services.RatesService) *View {
	return &View{Ctx: ctx, ratesService: ratesService}
}
