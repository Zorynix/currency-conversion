package routes

import (
	"context"
	"currency-conversion/services"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type RouterHead struct {
	MSQ  services.Database
	Addr *string
}

type Router struct {
	Router *fiber.App
	MSQ    services.Database
}

type Route struct {
	Group fiber.Router
	MSQ   services.Database
}

func Routes(addr *string) {

	mysql, err := services.NewMySQL(context.Background())
	if err != nil {
		log.Fatal().Err(err)
	}

	router := fiber.New()

	route := Router{Router: router, MSQ: mysql}

	route.V1Routes()

	log.Info().Msgf("Starting server on port %d...", 8000)
	if err := router.Listen(":8000"); err != nil {
		log.Fatal().Err(err).Msg("Can not start http server")
	}

}
