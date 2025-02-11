package main

import (
	"TanAgah/internal/config"
	"TanAgah/internal/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger.InitLogger()

}
