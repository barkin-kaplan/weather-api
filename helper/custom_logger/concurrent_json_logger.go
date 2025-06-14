package custom_logger

import (
	"encoding/json"
	"fmt"
	"time"
)

// LogLevel defines the severity of the log

// ConcurrentLogger struct to manage the log channel
type ConcurrentJsonLogger struct {
	logChannel chan string
	done       chan bool
	Name       string
	LogLevel   LogLevel
}

// NewConcurrentLogger creates a new logger instance
func NewConcurrentJsonLogger(name string, logLevel LogLevel) *ConcurrentJsonLogger {
	logger := &ConcurrentJsonLogger{
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
func (l *ConcurrentJsonLogger) logWorker() {
	for {
		select {
		case msg := <-l.logChannel:
			// Log message received, print to console
			fmt.Println(msg)
		case <-l.done:
			// Exit goroutine when done signal is received
			fmt.Print("Logger stopped %s", l.Name)
			return
		}
	}
}

func (l *ConcurrentJsonLogger) logMessage(level LogLevel, levelName string, message map[string]any) {
	// Check if the current log level allows this message to be logged
	if level < l.LogLevel {
		return
	}

	// Get the current timestamp with microsecond precision
	timestamp := time.Now().Format("2006-01-02 15:04:05.000000")
	message["timestamp"] = timestamp
	message["log_level"] = levelName
	message["log_name"] = l.Name

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error serializing JSON:", err)
		return
	}

	l.logChannel <- string(jsonBytes)
}

// Log methods for each log level

func (l *ConcurrentJsonLogger) DebugMap(message map[string]any) {
	l.logMessage(DEBUG, "DEBUG", message)
}

func (l *ConcurrentJsonLogger) InfoMap(message map[string]any) {
	l.logMessage(INFO, "INFO", message)
}

func (l *ConcurrentJsonLogger) WarnMap(message map[string]any) {
	l.logMessage(WARN, "WARN", message)
}

func (l *ConcurrentJsonLogger) ErrorMap(message map[string]any) {
	l.logMessage(ERROR, "ERROR", message)
}

func (l *ConcurrentJsonLogger) Debug(message string) {
	l.logMessage(DEBUG, "DEBUG", map[string]any{"message": message})
}

func (l *ConcurrentJsonLogger) Info(message string) {
	l.logMessage(INFO, "INFO", map[string]any{"message": message})
}

func (l *ConcurrentJsonLogger) Warn(message string) {
	l.logMessage(WARN, "WARN", map[string]any{"message": message})
}

func (l *ConcurrentJsonLogger) Error(message string) {
	l.logMessage(ERROR, "ERROR", map[string]any{"message": message})
}


// Stop stops the logger goroutine
func (l *ConcurrentJsonLogger) Stop() {
	l.done <- true
}
