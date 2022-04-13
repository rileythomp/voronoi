package handlers

import (
	"fmt"
	"log"
)

type Logger struct {
	log *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		log: log.Default(),
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.printf("INFO", format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printf("ERROR", format, v...)
}

func (l *Logger) printf(logType string, format string, v ...interface{}) {
	l.log.Printf(fmt.Sprintf("%5s: %s", logType, format), v...)
}
