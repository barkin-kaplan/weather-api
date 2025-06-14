package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/barkin-kaplan/weather-api/db"
	"github.com/barkin-kaplan/weather-api/integration"
	"github.com/barkin-kaplan/weather-api/server/model/resp"
)


type RequestGroup struct {
	location string
	requests []chan resp.WeatherResponse
	timer    *time.Timer
	mutex    sync.Mutex
	w http.ResponseWriter
	id int
}


type Server struct {
	requestGroups map[string]*RequestGroup
	groupsMutex   sync.Mutex
	weatherData integration.IWeatherData
	maxRequestCount int
	maxDelay time.Duration
	postgre db.IPostgreConnector
	idCounter int
}

func NewServer(weatherData integration.IWeatherData, maxRequestCount int, maxDelay time.Duration, postgre db.IPostgreConnector) *Server {
	return &Server{
		requestGroups: map[string]*RequestGroup{},
		weatherData: weatherData,
		maxRequestCount: maxRequestCount,
		maxDelay: maxDelay,
		postgre: postgre,
	}
}

func (s *Server) Start(port int) {
	s.routes()
	fmt.Printf("Starting server at port %d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

