package dorfyn

import (
	"log"
	"os"
)

// DorfynLogLevel represents the log level in use.
type DorfynLogLevel int

const (
	// LogNone disables logging.
	LogNone = 0
	// LogError logs only errors.
	LogError = 1
	// LogInfo logs errors and info messages.
	LogInfo = 2
	// LogDebug logs errors, info messages and debug messages.
	LogDebug = 3
)

var (
	// LogLevel is the log level to use. Defaults to LogNone.
	LogLevel DorfynLogLevel
	// Logger is the logger used by the Dorfyn library. Defaults to os.Stderr.
	Logger *log.Logger
)

// logError logs an error message. To be used when an error is returned from a remote call.
func logError(format string, v ...any) {
	if LogLevel >= LogError {
		Logger.Printf("[error] "+format, v...)
	}
}

// logInfo logs an info message. To be used when you want to log something that is not an error, but still might be relevant.
func logInfo(format string, v ...any) {
	if LogLevel >= LogInfo {
		Logger.Printf("[info] "+format, v...)
	}
}

// logDebug logs a debug message. To be used for more detailed logging.
func logDebug(format string, v ...any) {
	if LogLevel >= LogDebug {
		Logger.Printf("[debug] "+format, v...)
	}
}

func init() {
	LogLevel = LogNone
	Logger = log.New(os.Stderr, "", log.LstdFlags)
}
