package cep

import (
	"github.com/felipeivanaga/go-expert-weather-gcp/internal/internal_error"
)

type CepProvider interface {
    GetCityName(cep string) (string, *internal_error.InternalError)
}

