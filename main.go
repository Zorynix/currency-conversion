package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
	addr = flag.String("addr", ":8000", "TCP address to listen to")
)

func main() {
	//flag.Parse()

	//routes.Routes(addr)

	//currency_codes := []string{}

	// testDB()

	//fmt.Println(currency_codes)
	// create as
	//

	//services.TestApi()

	//flag.Parse()

	//routes.Routes(addr)

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error when uploading a file .env")
		return
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		client := &http.Client{}

		method := os.Getenv("methodGet")
		url := os.Getenv("url_all_currencies")

		req, err := http.NewRequest(os.Getenv(method), url, nil)
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

// for _, currencyCode := range currency_codes {
//go updateRate
// }

// fmt.Println("Hello, 世界")

func updateRates() {
	// Currencyapi Service # hourly
	// ECB Service # daily

	// rates Service interface

	// currency_codes := [180]string{"USD", "CAD", "CNY"}

	// for _, currencyCode := range currency_codes {
	// 	go updateRate
	// }
}

func updateRate( /* currencyCode */ ) {
	// get latest from API
	// update to DB
	// --- resave to history table

	// @TODO: add event to queue (rabbitmq)
}
