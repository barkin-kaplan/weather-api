package server

import (
	"fmt"
	"time"

	"github.com/barkin-kaplan/weather-api/db"
	"github.com/barkin-kaplan/weather-api/helper/slicehelper"
	"github.com/barkin-kaplan/weather-api/server/model/resp"
)

func (s *Server) registerToGroup(location string, channel chan resp.WeatherResponse) {
	s.groupsMutex.Lock()
	defer s.groupsMutex.Unlock()

	var group *RequestGroup
	groups, exists := s.requestGroups[location]
	if !exists {
		groups := []*RequestGroup{}
		s.groupIdCounter += 1
		group = &RequestGroup{
			location: location,
			requests: make([]chan resp.WeatherResponse, 0),
			groupId:  s.groupIdCounter,
		}
		groups = append(groups, group)
		s.requestGroups[location] = groups

		group.timer = time.AfterFunc(s.maxDelay, func() {
			s.process(group)
		})
	} else {

		// handle client batching
		lastGroup := groups[len(groups)-1]
		lastGroup.mutex.Lock()
		defer lastGroup.mutex.Unlock()
		if len(lastGroup.requests) == s.maxRequestCount {
			if lastGroup.timer.Stop() {
				go s.process(lastGroup)
			}
			s.logger.Debug(fmt.Sprintf("Creating new group because current is full"))
			s.groupIdCounter += 1
			group = &RequestGroup{
				location: location,
				requests: make([]chan resp.WeatherResponse, 0),
				groupId:  s.groupIdCounter,
			}
			groups = append(groups, group)
			s.requestGroups[location] = groups

			group.timer = time.AfterFunc(s.maxDelay, func() {
				s.process(group)
			})
		} else {
			group = lastGroup
		}
	}

	group.requests = append(group.requests, channel)
}

func (s *Server) process(g *RequestGroup) {
	s.logger.Debug(fmt.Sprintf("Processing group: %d", g.groupId))
	g.mutex.Lock()
	defer g.mutex.Unlock()

	service1Temp, err1 := s.weatherData.FetchWeatherAPI(g.location)
	service2Temp, err2 := s.weatherData.FetchWeatherStack(g.location)
	if err1 != nil || err2 != nil {
		s.logger.Error(fmt.Sprintf("Error occurred with weather apis err1: %s, err2: %s", err1, err2))
	}

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
	
	weatherQuery := &db.WeatherQuery{
		Location:            g.location,
		Service1Temperature: service1Temp,
		Service2Temperature: service2Temp,
		RequestCount:        len(g.requests),
	}
	s.postgre.SaveWeatherQuery(weatherQuery)
	// Clean up the request group
	s.groupsMutex.Lock()
	relatedGroups, exists := s.requestGroups[g.location]
	if exists && len(relatedGroups) == 1 {
		delete(s.requestGroups, g.location)
	} else {
		slicehelper.FindAndRemove(relatedGroups, func(group *RequestGroup) bool {
			return group.groupId == g.groupId
		})
	}

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
