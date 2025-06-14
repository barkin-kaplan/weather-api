package server

import "net/http"
func (s * Server) routes() {
	http.HandleFunc("/weather", s.weather)
}