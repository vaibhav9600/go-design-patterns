package logger

import (
	"encoding/json"
	"fmt"
)

// json logger logs messages in json struct
type JSONLogger struct{}

func (j *JSONLogger) LogJSON(message string) {
	log := map[string]string{"message": message}
	jsonLog, _ := json.Marshal(log)
	fmt.Println(string(jsonLog))
}

type JSONLoggerAdapter struct {
	jsonLogger *JSONLogger
}

func NewJsonLoggerAdapter(jsonLogger *JSONLogger) *JSONLoggerAdapter {
	return &JSONLoggerAdapter{jsonLogger: jsonLogger}
}

func (adapter *JSONLoggerAdapter) Log(message string) {
	adapter.jsonLogger.LogJSON(message)
}
