package main

import (
	"context"
	"currency-conversion/config"
	"currency-conversion/routes"
	"currency-conversion/services"
	"currency-conversion/utils"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var (
	addr       = flag.String("addr", ":8000", "TCP address to listen to")
	configPath = flag.String("config", "config/config.yaml", "Path to config file")
)

func main() {
	flag.Parse()

	err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	utils.InitLogger(config.Cfg.LogLevel)

	mysqlService, err := services.NewMySQL(context.Background(), config.Cfg.DSN)
	if err != nil {
		logrus.Fatalf("Failed to initialize MySQL service: %v", err)
	}

	app := fiber.New()

	router := routes.NewRouter(app, mysqlService)
	router.SetupRoutes()

	logrus.Infof("Starting server on %s...", *addr)
	if err := app.Listen(*addr); err != nil {
		logrus.Fatalf("Failed to start HTTP server: %v", err)
	}
}
