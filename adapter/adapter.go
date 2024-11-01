package main

import "adapter/logger"

func main() {
	jsonLogger := logger.JSONLogger{}

	adapter := logger.NewJsonLoggerAdapter(&jsonLogger)

	adapter.Log("Hello, World!")

	logger.Process(adapter, "hello, world!")
}
