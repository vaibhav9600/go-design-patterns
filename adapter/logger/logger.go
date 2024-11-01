package logger

// Logger defines the interface with a simple Log method
type Logger interface {
	Log(message string)
}

// Process Logs a message using any struct that implements Logger interface
func Process(l Logger, message string) {
	l.Log(message)
}
