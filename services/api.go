package services

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic().Msg("No .env file found")
	}
}

var (
	addr                      = flag.String("addr", ":8000", "TCP address to listen to")
	url_all_currencies        = "https://api.currencyapi.com/v3/currencies"
	url_latest_exchange_rates = "https://api.currencyapi.com/v3/latest"
	methodGet                 = "GET"
)

func TestApi() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error when uploading a file .env")
		return
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		client := &http.Client{}

		req, err := http.NewRequest(methodGet, url_all_currencies, nil)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error when creating a request: %s", err))
		}

		apiKEY := os.Getenv("API_KEY")

		req.Header.Add("apikey", apiKEY)

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
	app.Listen(*addr)
}
