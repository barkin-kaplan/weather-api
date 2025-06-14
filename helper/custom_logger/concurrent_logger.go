package custom_logger

import (
	"fmt"
	"time"
)

// LogLevel defines the severity of the log

// ConcurrentLogger struct to manage the log channel
type ConcurrentLogger struct {
	logChannel chan string
	done       chan bool
	Name       string
	LogLevel   LogLevel
}

// NewConcurrentLogger creates a new logger instance
func NewConcurrentLogger(name string, logLevel LogLevel) *ConcurrentLogger {
	logger := &ConcurrentLogger{
		logChannel: make(chan string, 5000), // Buffered channel to hold logs
		done:       make(chan bool),
		Name:       name,
		LogLevel:   logLevel,
	}

	// Start the logger in a separate goroutine
	go logger.logWorker()

	return logger
}

// logWorker is a goroutine that processes log messages
func (l *ConcurrentLogger) logWorker() {
	for {
		select {
		case msg := <-l.logChannel:
			// Log message received, print to console
			fmt.Print(msg)
		case <-l.done:
			// Exit goroutine when done signal is received
			fmt.Printf("Logger stopped %s", l.Name)
			return
		}
	}
}

func (l *ConcurrentLogger) logMessage(level LogLevel, levelName string, message string) {
	// Check if the current log level allows this message to be logged
	if level < l.LogLevel {
		return
	}

	// Get the current timestamp with microsecond precision
	timestamp := time.Now().Format("2006-01-02 15:04:05.000000")

	// Print the log message in the format: [timestamp] [level] message
	msg := fmt.Sprintf("%s [%s] [%s] %s\n", timestamp, l.Name, levelName, message)
	l.logChannel <- msg
}

// Log methods for each log level

func (l *ConcurrentLogger) Debug(message string) {
	l.logMessage(DEBUG, "DEBUG", message)
}

func (l *ConcurrentLogger) Info(message string) {
	l.logMessage(INFO, "INFO", message)
}

func (l *ConcurrentLogger) Warn(message string) {
	l.logMessage(WARN, "WARN", message)
}

func (l *ConcurrentLogger) Error(message string) {
	l.logMessage(ERROR, "ERROR", message)
}

// Stop stops the logger goroutine
func (l *ConcurrentLogger) Stop() {
	l.done <- true
}
