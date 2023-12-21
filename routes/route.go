package routes

import (
	"context"
	"currency-conversion/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type RouterHead struct {
	MSQ  *services.Mysql
	Addr *string
}

type Router struct {
	Router *fiber.App
	MSQ    *services.Mysql
}

type Route struct {
	Group fiber.Router
	MSQ   *services.Mysql
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic().Msg("No .env file found")
	}
}

func Routes(addr *string) {
	mysql, err := services.NewMySQL(context.Background())
	if err != nil {
		log.Fatal().Err(err)
	}

	router := fiber.New()

	route := Router{Router: router, MSQ: mysql}

	route.V1Routes()

	router.Get("/", func(c *fiber.Ctx) error {
		url := os.Getenv("url_all_currencies")
		method := os.Getenv("method")
		client := &http.Client{}

		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error when creating a request: %s", err))
		}

		apiKey := os.Getenv("API_KEY")
		req.Header.Add("apikey", apiKey)

		res, err := client.Do(req)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error when executing a request: %s", err))
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error reading data from the response: %s", err))
		}

		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error during JSON parsing: %s", err))
		}

		prettiedJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error when formatting JSON: %s", err))
		}

		return c.SendString(string(prettiedJSON))
	})

	if err := router.Listen(*addr); err != nil {
		log.Fatal().Err(err).Msg("Can not start http server")
	}
}
