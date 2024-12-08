// pkg/logger/logger.go
package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

func NewLogger(level string, output io.Writer) (*logrus.Entry, error) {
	baseLogger := logrus.New()

	if output != nil {
		baseLogger.SetOutput(output)
	}

	baseLogger.SetFormatter(&logrus.JSONFormatter{})

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %v", err)
	}

	baseLogger.SetLevel(logLevel)

	return baseLogger.WithFields(logrus.Fields{}), nil
}

func WithFields(log *logrus.Entry, fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}
