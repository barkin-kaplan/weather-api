package integration

type WeatherStackResponse struct {
	Current struct {
		Temperature float64 `json:"temperature"`
	} `json:"current"`
}