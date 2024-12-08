package config

import (
	"fmt"
	"github.com/jeffery/logger/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
)

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Fprintf(os.Stderr, "%s environment variable must be set\n", key)
		os.Exit(1)
	}
	return value
}

func InitializeLogger(logLevel string, envVarFields map[string]string) *logrus.Entry {
	log, err := logger.NewLogger(logLevel, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating logger: %v\n", err)
		os.Exit(1)
	}

	fields := logrus.Fields{}
	for fieldName, envVar := range envVarFields {
		fields[fieldName] = getRequiredEnv(envVar)
	}

	return log.WithFields(fields)
}
