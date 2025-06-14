package req

type WeatherRequest struct {
	Query string `json="q"`
}