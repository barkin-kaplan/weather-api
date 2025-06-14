package custom_logger

import (
	"fmt"
	"time"
)


// CustomLogger struct to hold the current log level
type CustomLogger struct {
	Name string
	LogLevel LogLevel
}

// logMessage logs a message with a timestamp and log level
func (l *CustomLogger) logMessage(level LogLevel, levelName string, message string) {
	// Check if the current log level allows this message to be logged
	if level < l.LogLevel {
		return
	}

	// Get the current timestamp with microsecond precision
	timestamp := time.Now().Format("2006-01-02 15:04:05.000000")

	// Print the log message in the format: [timestamp] [level] message
	fmt.Printf("%s [%s] [%s] %s\n", timestamp, l.Name, levelName, message)
}

// Log methods for each log level

func (l *CustomLogger) Debug(message string) {
	l.logMessage(DEBUG, "DEBUG", message)
}

func (l *CustomLogger) Info(message string) {
	l.logMessage(INFO, "INFO", message)
}

func (l *CustomLogger) Warn(message string) {
	l.logMessage(WARN, "WARN", message)
}

func (l *CustomLogger) Error(message string) {
	l.logMessage(ERROR, "ERROR", message)
}

