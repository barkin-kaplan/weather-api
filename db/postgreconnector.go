package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/barkin-kaplan/weather-api/helper/custom_logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IPostgreConnector interface {
	SaveWeatherQuery(query *WeatherQuery)
}

type PostgreConnector struct {
	db      *gorm.DB
	timeout time.Duration
	logger  *custom_logger.ConcurrentLogger
}

func NewPostgreConnector(connString string, timeout time.Duration) *PostgreConnector {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return &PostgreConnector{
		db:      db,
		timeout: timeout,
		logger:  custom_logger.NewConcurrentLogger("PostgreConnector", custom_logger.DEBUG),
	}
}

func (p *PostgreConnector) Migrate() error {
	return p.db.AutoMigrate(&WeatherQuery{})
}


func (p *PostgreConnector) SaveWeatherQuery(query *WeatherQuery) {
	go func(query *WeatherQuery) {
		newCtx, cancel := context.WithTimeout(context.Background(), p.timeout)
		defer cancel()
		err := p.db.WithContext(newCtx).Create(query).Error
		if err != nil {
			p.logger.Error(fmt.Sprintf("Error occurred writing query to db %s", err))
		}
	}(query)

}


type MockPostgreConnector struct {

}

func (p * MockPostgreConnector) SaveWeatherQuery(query *WeatherQuery) {

}