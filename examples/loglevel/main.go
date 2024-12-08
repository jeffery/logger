package main

import (
	"github.com/jeffery/logger/pkg/config"
	"github.com/jeffery/logger/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	envVarFields := map[string]string{
		"env":     "ENV",
		"service": "SERVICE",
		"version": "VERSION",
	}

	log := config.InitializeLogger(os.Getenv("LOG_LEVEL"), envVarFields)

	// Add service-specific correlation ID if needed
	log = logger.WithFields(log, logrus.Fields{
		"correlationId": "123-456-789",
	})

	log.Trace("This is a trace message")
	log.Debug("This is a debug message")
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")
}
