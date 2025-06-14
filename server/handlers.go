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

	s.idCounter += 1
	id := s.idCounter
	fmt.Printf("Request with id %d arrived\n", id)
    responseChan := make(chan resp.WeatherResponse, 1) // <--- buffered to prevent blocking
    group := s.getOrCreateRequestGroup(location, w)
    s.addRequest(responseChan, group)

    select {
    case response := <-responseChan:
		fmt.Printf("Request with id %d served\n", id)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    case <-time.After(10 * time.Second):  // safety timeout
        http.Error(w, "Request timeout", http.StatusGatewayTimeout)
    }
}
