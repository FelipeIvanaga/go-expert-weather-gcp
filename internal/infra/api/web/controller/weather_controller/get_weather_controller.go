package weathercontroller

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/felipeivanaga/go-expert-weather-gcp/internal/configuration/rest_err"
	weatherusecase "github.com/felipeivanaga/go-expert-weather-gcp/internal/usecase/weather_usecase"
)

type WeatherController struct {
	weatherUsecase weatherusecase.WeatherCaseInterface
}

func NewWeatherController(weatherUsecase weatherusecase.WeatherCaseInterface) *WeatherController {
	return &WeatherController{
		weatherUsecase: weatherUsecase,
	}
}

func (w *WeatherController) GetWeather(ctx *gin.Context) {
	cep := ctx.Query("CEP")

	re := regexp.MustCompile("^[0-9]{8}$")

	if !re.MatchString(cep) {
		errRest := rest_err.NewUnprocessableEntityError("invalid zipcode")
		ctx.JSON(errRest.Code, errRest)
		return
	}

	weatherData, err := w.weatherUsecase.GetWeather(ctx, cep)
	if err != nil {
		restErr := rest_err.ConvertError(err)

		ctx.JSON(restErr.Code, restErr)
		return
	}

	ctx.JSON(http.StatusOK, weatherData)
}
