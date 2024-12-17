package weatherusecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/felipeivanaga/go-expert-weather-gcp/internal/infra/provider/weather"
	"github.com/felipeivanaga/go-expert-weather-gcp/internal/internal_error"
)

type MockCepProvider struct {
	mock.Mock
}

func (m *MockCepProvider) GetCityName(cep string) (string, *internal_error.InternalError) {
	args := m.Called(cep)
	if err := args.Get(1); err != nil {
		return args.String(0), err.(*internal_error.InternalError) // Type assertion should be safe
	}
	return args.String(0), nil
}

// MockWeatherProvider is a mock implementation of the WeatherProvider interface
type MockWeatherProvider struct {
	mock.Mock
}

func (m *MockWeatherProvider) GetWeatherWithCityName(cityName string) (*weather.GetWeatherResponseDTO, *internal_error.InternalError) {
	args := m.Called(cityName)
	if err := args.Get(1); err != nil {
		return nil, err.(*internal_error.InternalError) // Type assertion should be safe
	}
	return args.Get(0).(*weather.GetWeatherResponseDTO), nil
}

func TestGetWeather_Success(t *testing.T) {
	// Create mocks for CepProvider and WeatherProvider
	mockCepProvider := new(MockCepProvider)
	mockWeatherProvider := new(MockWeatherProvider)

	// Expected city name and weather data
	expectedCity := "CityName"
	expectedWeather := &weather.GetWeatherResponseDTO{
		Celsius:    25.0,
		Fahrenheit: 77.0,
	}

	// Define mock behavior for CepProvider and WeatherProvider
	mockCepProvider.On("GetCityName", "12345678").Return(expectedCity, nil)
	mockWeatherProvider.On("GetWeatherWithCityName", expectedCity).Return(expectedWeather, nil)

	// Create the WeatherUsecase
	weatherUsecase := NewWeatherUsecase(mockCepProvider, mockWeatherProvider)

	// Call the method to be tested
	result, err := weatherUsecase.GetWeather(context.Background(), "12345678")

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedWeather.Celsius, result.TempC)
	assert.Equal(t, expectedWeather.Celsius+273, result.TempK)
	assert.Equal(t, expectedWeather.Fahrenheit, result.TempF)

	// Verify that the mock methods were called
	mockCepProvider.AssertExpectations(t)
	mockWeatherProvider.AssertExpectations(t)
}

func TestGetWeather_CepProviderError(t *testing.T) {
	// Create mocks for CepProvider and WeatherProvider
	mockCepProvider := new(MockCepProvider)
	mockWeatherProvider := new(MockWeatherProvider)

	// Define mock behavior for CepProvider to return an error
	mockCepProvider.On("GetCityName", "12345678").Return("", internal_error.NewInternalServerError("failed to get city"))

	// Create the WeatherUsecase
	weatherUsecase := NewWeatherUsecase(mockCepProvider, mockWeatherProvider)

	// Call the method to be tested
	result, err := weatherUsecase.GetWeather(context.Background(), "12345678")

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to get city", err.Message)

	// Verify that the mock methods were called
	mockCepProvider.AssertExpectations(t)
}

func TestGetWeather_WeatherProviderError(t *testing.T) {
	// Create mocks for CepProvider and WeatherProvider
	mockCepProvider := new(MockCepProvider)
	mockWeatherProvider := new(MockWeatherProvider)

	// Expected city name
	expectedCity := "CityName"

	// Define mock behavior for CepProvider and WeatherProvider
	mockCepProvider.On("GetCityName", "12345678").Return(expectedCity, nil)
	mockWeatherProvider.On("GetWeatherWithCityName", expectedCity).Return(nil, internal_error.NewInternalServerError("failed to get weather"))

	// Create the WeatherUsecase
	weatherUsecase := NewWeatherUsecase(mockCepProvider, mockWeatherProvider)

	// Call the method to be tested
	result, err := weatherUsecase.GetWeather(context.Background(), "12345678")

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to get weather", err.Message)

	// Verify that the mock methods were called
	mockCepProvider.AssertExpectations(t)
	mockWeatherProvider.AssertExpectations(t)
}
