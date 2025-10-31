package log

import (
	"campushelphub/internal/errors"
	"log"
)

type Logger struct {
	Log *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Log: log.New(log.Writer(), "CampusHelpHub: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(data interface{}) {
	l.Log.Println("INFO:", data)
}

func (l *Logger) Error(err *errors.Error) {
	l.Log.Println("ERROR:", err.Error())
}
