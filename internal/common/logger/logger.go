package logger

import (
	"strings"

	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
	"github.com/sirupsen/logrus"
)


type logrusLogger struct {
	entry *logrus.Entry
}

func (ll *logrusLogger) Error(args ...interface{}) {
	ll.entry.Error(args...)
}
func (ll *logrusLogger) Debug(args ...interface{}) {
	ll.entry.Debug(args...)
}
func (ll *logrusLogger) Info(args ...interface{}) {
	ll.entry.Info(args...)
}
func (ll *logrusLogger) Warn(args ...interface{}) {
	ll.entry.Warn(args...)
}
func (ll *logrusLogger) Panic(args ...interface{}) {
	ll.entry.Panic(args...)
}
func (ll *logrusLogger) WithFields(fields map[string]interface{}) ports.Logger {
	return &logrusLogger{entry: ll.entry.WithFields(fields)}
}

func (ll *logrusLogger) WithField(key string, value interface{}) ports.Logger {
	return &logrusLogger{entry: ll.entry.WithField(key, value)}
}

func NewLogger(logLevel string) ports.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		 level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	return &logrusLogger{entry: logrus.NewEntry(logger)}
}