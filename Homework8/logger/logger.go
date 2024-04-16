package logger

import (
	"Hillel/Homework8/config"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(cfg *config.Configuration) *Logger {
	l := logrus.New()

	parsedLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Info("failed to parse log level. log level will be set [info]")
		parsedLevel = logrus.InfoLevel
	}

	l.SetLevel(parsedLevel)

	return &Logger{
		Logger: l,
	}
}
