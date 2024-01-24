package services

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type LoggerImp struct {
	log *logrus.Logger
}

func NewLogger(level string) Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Panicf("Invalid log level: %s, using default level (Info)", level)
	}
	log.SetLevel(logLevel)
	return &LoggerImp{
		log: log,
	}
}

func (l *LoggerImp) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *LoggerImp) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *LoggerImp) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l *LoggerImp) Panic(args ...interface{}) {
	l.log.Panic(args...)
}
