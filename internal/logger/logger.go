package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(level string) *logrus.Logger {
	log := logrus.New()
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.WarnLevel)
	}
	return log
}
