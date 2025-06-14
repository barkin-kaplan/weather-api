package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/barkin-kaplan/weather-api/server/model/resp"
)

func (s *Server) weather(w http.ResponseWriter, r *http.Request){
	location := r.URL.Query().Get("q")
    if location == "" {
        http.Error(w, "Missing 'q' query param", http.StatusBadRequest)
        return
    }

	s.requestIdCounter += 1
	id := s.requestIdCounter
	s.logger.Info(fmt.Sprintf("Request with id %d arrived", id))
    responseChan := make(chan resp.WeatherResponse, 1) // <--- buffered to prevent blocking
    s.registerToGroup(location, responseChan)

    select {
    case response := <-responseChan:
		s.logger.Info(fmt.Sprintf("Request with id %d served", id))
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    case <-time.After(10 * time.Second):  // safety timeout
        http.Error(w, "Request timeout", http.StatusGatewayTimeout)
    }
}
