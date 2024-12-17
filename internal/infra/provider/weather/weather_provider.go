package weather

import "github.com/felipeivanaga/go-expert-weather-gcp/internal/internal_error"

type GetWeatherResponseDTO struct {
    Celsius float64
    Fahrenheit float64
}

type WeatherProvider interface {
    GetWeatherWithCityName(city string) (*GetWeatherResponseDTO, *internal_error.InternalError)
}

