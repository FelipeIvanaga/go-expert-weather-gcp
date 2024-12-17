package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	weathercontroller "github.com/felipeivanaga/go-expert-weather-gcp/internal/infra/api/web/controller/weather_controller"
	"github.com/felipeivanaga/go-expert-weather-gcp/internal/infra/provider/cep"
	"github.com/felipeivanaga/go-expert-weather-gcp/internal/infra/provider/weather"
	weatherusecase "github.com/felipeivanaga/go-expert-weather-gcp/internal/usecase/weather_usecase"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	router := gin.Default()

	wheaterController := initDependencies()

	router.GET("/weather", wheaterController.GetWeather)

	router.Run(getEnvPort())
}

func initDependencies() (weatherController *weathercontroller.WeatherController) {
	weatherUsecase := weatherusecase.NewWeatherUsecase(cep.NewViaCepProvider(), weather.NewWeatherapiProvider(getWeatherApiToken()))
	weatherController = weathercontroller.NewWeatherController(weatherUsecase)
	return
}

func getEnvPort() string {
	return ":" + os.Getenv("PORT")
}

func getWeatherApiToken() string {
	return os.Getenv("TOKEN_WEATHERAPI")
}
