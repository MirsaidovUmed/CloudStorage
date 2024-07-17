package logger

import (
	"CloudStorage/pkg/config"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(config *config.Config) *logrus.Logger {
	f, err := os.OpenFile(config.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o777)
	if err != nil {
		panic(err)
	}

	log := &logrus.Logger{
		Out:       io.MultiWriter(f, os.Stdout),
		Level:     logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{},
	}
	log.SetReportCaller(true)

	return log
}
