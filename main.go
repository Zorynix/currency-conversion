package main

import (
	"currency-conversion/routes"
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

// func testDB() error {

// 	err := godotenv.Load()

// 	user := os.Getenv("user")
// 	password := os.Getenv("password")
// 	host := os.Getenv("host")
// 	port := os.Getenv("port")
// 	dbname := os.Getenv("dbname")

// 	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname))

// 	if err != nil {
// 		fmt.Println("error validating sql.Open arguments")
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	err = db.Ping()

// 	if err != nil {
// 		fmt.Println("error verifying connection with db.Ping")
// 		panic(err.Error())
// 	}

// 	// insert, err := db.Query("INSERT INTO `razzzila`.`currencies_exchange_rates` (`id`,`currency_id`,`target_currency_id`,`exchange_rate`,`rate_source_id`,`created_at`,`updated_at`) VALUES ('9','99','999','3','9','2024-12-12 23:24:25','2025-12-13 22:23:24');")
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	// defer insert.Close()

// 	// fmt.Println("Successful insertion")
// 	return nil
// }

var (
	addr = flag.String("addr", ":8000", "TCP address to listen to")
)

func main() {

	routes.Routes(addr)

	currency_codes := []string{}

	// testDB()

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error when uploading a file .env")
		return
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

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
	app.Listen(":8080")
	fmt.Println(currency_codes)
	// create as
	//

	// for _, currencyCode := range currency_codes {
	//go updateRate
	// }

	// fmt.Println("Hello, 世界")
}

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
