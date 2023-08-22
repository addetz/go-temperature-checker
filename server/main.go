package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/addetz/railway-go-demo/server/weather"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	TIMEOUT          = 3 * time.Second
	WEATHER_BASE_URL = "http://api.openweathermap.org/data/2.5/weather"
	KELVIN_CONSTANT  = 273.15
)

func main() {
	// Read port if one is set
	port := readPort()

	// Read API key
	apiKey := readWeatherAPIKey()

	// Initialise echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configure server
	s := http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           e,
		ReadTimeout:       TIMEOUT,
		ReadHeaderTimeout: TIMEOUT,
		WriteTimeout:      TIMEOUT,
		IdleTimeout:       TIMEOUT,
	}

	// Set up the root route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Hello, Railway Gophers!",
		})
	})

	e.GET("/weather/:city", func(c echo.Context) error {
		city := c.Param("city")
		weatherURL := getWeatherURL(apiKey, city)
		log.Println(weatherURL)
		response, err := http.Get(weatherURL)
		if err != nil {
			return echo.NewHTTPError(response.StatusCode, err)
		}
		jsonResponse := new(weather.WeatherResponse)
		if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message":     fmt.Sprintf("Fetching data for %s", city),
			"feels_like":  convertCelsius(jsonResponse.Main.FeelsLike),
			"temp":        convertCelsius(jsonResponse.Main.Temp),
			"description": jsonResponse.Weather[0].Description,
		})
	})

	log.Printf("Listening on :%s...\n", port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}

// readPort reads the SERVER_PORT environment variable if one is set
// or returns a default if none is found
func readPort() string {
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		return "1323"
	}
	return port
}

// readWeatherAPIKey reads the readWeatherAPIKey environment variable if one is set
// or fatally exits the server if it is not set
func readWeatherAPIKey() string {
	key, ok := os.LookupEnv("WEATHER_API_KEY")
	if !ok {
		log.Fatal("WEATHER_API_KEY must be set")
	}
	return key
}

// getWeatherURL constructs the weather url based on city and API key
func getWeatherURL(apiKey, city string) string {
	return fmt.Sprintf("%s?q=%s&appid=%s", WEATHER_BASE_URL, city, apiKey)
}

// convertCelsius takes a Kelvin temp and returns the Celsius number
func convertCelsius(kTemp float64) string {
	return fmt.Sprintf("%.2f", kTemp-KELVIN_CONSTANT)
}
