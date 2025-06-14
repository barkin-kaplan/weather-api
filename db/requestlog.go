package db

import (
    "time"
)

type WeatherQuery struct {
    ID                  uint           `gorm:"primaryKey;autoIncrement;column:id"`
    Location            string         `gorm:"type:text;not null;column:location"`
    Service1Temperature float64        `gorm:"type:double precision;column:service_1_temperature"`
    Service2Temperature float64        `gorm:"type:double precision;column:service_2_temperature"`
    RequestCount        int            `gorm:"column:request_count"`
    CreatedAt           time.Time      `gorm:"autoCreateTime;column:created_at"`
}

// Explicit table name (optional, but good practice for PostgreSQL)
func (WeatherQuery) TableName() string {
    return "weather_queries"
}
