package app

import (
	"context"
	"currency-conversion/config"
	"currency-conversion/internal/routes"
	"currency-conversion/internal/services"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Run(configPath string) {
	err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}
	SetLogrus(config.Cfg.LogLevel)

	mysqlService, err := services.NewMySQL(context.Background(), config.Cfg.DSN)
	if err != nil {
		logrus.Fatalf("Failed to initialize MySQL service: %v", err)
	}
	app := fiber.New()

	router := routes.NewRouter(app, mysqlService)
	router.SetupRoutes()

	addr := config.Cfg.Server.Address
	logrus.Infof("Starting server on %s...", addr)
	go func() {
		if err := app.Listen(addr); err != nil {
			logrus.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	logrus.Info("Configuring graceful shutdown...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logrus.Info("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		logrus.Fatalf("Server forced to shutdown: %v", err)
	}

	logrus.Info("Server exiting")
}
