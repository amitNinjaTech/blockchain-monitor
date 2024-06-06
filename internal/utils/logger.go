package utils

import (
	"log"
)

type Logger struct {
    logLevel string
}

func NewLogger(level string) *Logger {
    return &Logger{
        logLevel: level,
    }
}

func (l *Logger) Info(format string, v ...interface{}) {
    if l.logLevel == "info" || l.logLevel == "debug" {
        log.Printf("[INFO] "+format, v...)
    }
}

func (l *Logger) Error(format string, v ...interface{}) {
    log.Printf("[ERROR] "+format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
    if l.logLevel == "debug" {
        log.Printf("[DEBUG] "+format, v...)
    }
}
