package integration

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/barkin-kaplan/weather-api/server/model/integration"
)

type IWeatherData interface {
	FetchWeatherAPI(location string) (float64, error)
	FetchWeatherStack(location string) (float64, error)
}

type WeatherData struct {
	weatherApiKey   string
	weatherApiUrl   string
	weatherStackKey string
	weatherStackUrl string
}

func NewWeatherData(weatherApiKey, weatherApiUrl, weatherStackKey, weatherStackUrl string) *WeatherData {
	return &WeatherData{
		weatherApiKey:   weatherApiKey,
		weatherApiUrl:   weatherApiUrl,
		weatherStackKey: weatherStackKey,
		weatherStackUrl: weatherStackUrl,
	}
}

func (s *WeatherData) FetchWeatherAPI(location string) (float64, error) {
	url := fmt.Sprintf("%s?key=%s&q=%s", s.weatherApiUrl, s.weatherApiKey, location)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var apiResp integration.WeatherAPIResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return 0, err
	}
	if apiResp.Current.TempC == 0 {
		return 0, errors.New("wrong api format")
	}
	return apiResp.Current.TempC, nil
}

func (s *WeatherData) FetchWeatherStack(location string) (float64, error) {
	url := fmt.Sprintf("%s?access_key=%s&query=%s", s.weatherStackUrl, s.weatherStackKey, location)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var stackResp integration.WeatherStackResponse
	err = json.Unmarshal(body, &stackResp)
	if err != nil {
		return 0, err
	}
	if stackResp.Current.Temperature == 0 {
		return 0, errors.New("wrong api format")
	}
	return stackResp.Current.Temperature, nil
}

type MockWeatherData struct {
}

func (s *MockWeatherData) FetchWeatherAPI(location string) (float64, error) {
	return 24, nil
}

func (s *MockWeatherData) FetchWeatherStack(location string) (float64, error) {
	return 24, nil
}
