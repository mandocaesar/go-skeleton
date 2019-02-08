package utility

import (
	"github.com/mandocaesar/go-skeleton/config"
	"github.com/sebest/logrusly"
	"github.com/sirupsen/logrus"
)

//Log : struct for logger DI
type Log struct {
	logger *logrus.Logger
}

//NewLogger : instantiate new logger
func NewLogger(cfg config.Configuration) (*Log, error) {
	log := logrus.New()
	config := cfg.Loggly

	log.WithFields(logrus.Fields{
		"service": "Inem",
		"version": "0.1",
	})
	hook := logrusly.NewLogglyHook(config.Token, config.Host, logrus.WarnLevel, config.Tags...)
	log.Hooks.Add(hook)

	return &Log{logger: log}, nil
}

//LogInfo : log as info
func (l *Log) LogInfo(data interface{}) {
	l.logger.Info(data)
}

//LogError : log as error
func (l *Log) LogError(data interface{}) {
	l.logger.Info(data)
}
