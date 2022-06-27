package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogRus() *logrus.Logger {
	log := logrus.New()
	log.Out = os.Stdout

	return log
}

type Logging struct {
	logger Logger
}

func NewLogging(logger Logger) *Logging {
	return &Logging{logger: logger}
}

func (l *Logging) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logging) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logging) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}
