package log

import (
	"campushelphub/internal/errors"
	"log"
)

type BusinessLogInfo struct {
	BusinessType string
	ID           uint64
	Status       string
	Extra        map[string]interface{}
}

type Logger struct {
	Log *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Log: log.New(log.Writer(), "CampusHelpHub: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(data *BusinessLogInfo) {
	l.Log.Printf("INFO: business_type=%s, id=%d, status=%s, extra=%v", data.BusinessType, data.ID, data.Status, data.Extra)
}

func (l *Logger) Error(err *errors.Error) {
	l.Log.Println("ERROR:", err.Error())
}
