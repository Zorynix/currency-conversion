package main

import (
	"context"
	"currency-conversion/routes"
	"currency-conversion/services"
	"currency-conversion/utils"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

var (
	addr = flag.String("addr", ":8000", "TCP address to listen to")
)

func main() {
	flag.Parse()

	utils.InitLogger()

	mysqlService, err := services.NewMySQL(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize MySQL service")
	}

	app := fiber.New()

	router := routes.NewRouter(app, mysqlService)
	router.SetupRoutes()

	log.Info().Msgf("Starting server on %s...", *addr)
	if err := app.Listen(*addr); err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
