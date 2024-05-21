package logger

import "github.com/sirupsen/logrus"

// Log struct logger
type Log struct {
	log *logrus.Logger
}

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

// New init logger
func New() *Log {
	return &Log{
		log: logrus.New(),
	}
}

func (l *Log) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *Log) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *Log) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *Log) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}
