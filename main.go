package main

import (
	"log"

	"time"

	"github.com/barkin-kaplan/weather-api/db"
	"github.com/barkin-kaplan/weather-api/helper"
	"github.com/barkin-kaplan/weather-api/integration"
	"github.com/barkin-kaplan/weather-api/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error occurred loading .env file")
	}
	postgreConnString := helper.CheckAndGetEnvString("POSTGRE_CONN_STRING")
	postgreConnector := db.NewPostgreConnector(postgreConnString, time.Second * 5)
	postgreConnector.Migrate()
	weatherApiKey := helper.CheckAndGetEnvString("WEATHER_API_KEY")
	weatherApiUrl := helper.CheckAndGetEnvString("WEATHER_API_URL")
	weatherStackKey := helper.CheckAndGetEnvString("WEATHER_STACK_KEY")
	weatherStackUrl := helper.CheckAndGetEnvString("WEATHER_STACK_URL")
	maxRequestCount, err := helper.CheckAndGetEnvInteger("MAX_REQUEST_COUNT")
	if err != nil {
		log.Fatalf("Invalid MAX_REQUEST_COUNT parameter")
	}
	maxDelayRaw, err := helper.CheckAndGetEnvInteger("MAX_DELAY_SECONDS")
	if err != nil {return}
	serverPort, err := helper.CheckAndGetEnvInteger("SERVER_PORT")
	if err != nil {return}

	weatherData := integration.NewWeatherData(
		weatherApiKey,
		weatherApiUrl,
		weatherStackKey, 
		weatherStackUrl, 
	)
	server := server.NewServer(
		weatherData,
		maxRequestCount, 
		time.Duration(maxDelayRaw)*time.Second, 
		postgreConnector,
	)

	server.Start(serverPort)

}
