package server

import (
	"net/http"
	"time"

	"github.com/barkin-kaplan/weather-api/db"
	"github.com/barkin-kaplan/weather-api/server/model/resp"
)

func (s *Server) getOrCreateRequestGroup(location string, w http.ResponseWriter) *RequestGroup {
	s.groupsMutex.Lock()
	defer s.groupsMutex.Unlock()

	group, exists := s.requestGroups[location]
	if !exists {
		group = &RequestGroup{
			location: location,
			requests: make([]chan resp.WeatherResponse, 0),
			w: w,
		}
		s.requestGroups[location] = group

		group.timer = time.AfterFunc(s.maxDelay, func() {
			s.process(group)
		})
	}
	return group
}

func (s *Server) process(g *RequestGroup) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	service1Temp, err1 := s.weatherData.FetchWeatherAPI(g.location)
	service2Temp, err2 := s.weatherData.FetchWeatherStack(g.location)

	var avgTemp float64
	if err1 == nil && err2 == nil {
		avgTemp = (service1Temp + service2Temp) / 2
	} else if err1 == nil {
		avgTemp = service1Temp
	} else if err2 == nil {
		avgTemp = service2Temp
	} else {
		avgTemp = 0 // fallback if both fail
	}

	response := resp.WeatherResponse{
		Location:    g.location,
		Temperature: avgTemp,
	}

	for _, ch := range g.requests {
		ch <- response
	}
	weatherQuery := &db.WeatherQuery {
		Location: g.location,
		Service1Temperature: service1Temp,
		Service2Temperature: service2Temp,
		RequestCount: len(g.requests),
	}
	s.postgre.SaveWeatherQuery(weatherQuery)
	// Clean up the request group
	s.groupsMutex.Lock()
	delete(s.requestGroups, g.location)
	s.groupsMutex.Unlock()
}

func (s *Server) addRequest(ch chan resp.WeatherResponse, g *RequestGroup) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.requests = append(g.requests, ch)

	if len(g.requests) >= s.maxRequestCount {
		if g.timer.Stop() {
			go s.process(g)
		}
	}
}