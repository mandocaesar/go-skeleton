package utility

import (
	"github.com/Sirupsen/logrus"
)

//Log : struct for logger DI
type Log struct {
	logger *logrus.Logger
}

//NewLogger : instantiate new logger
func NewLogger() (*Log, error) {
	log := logrus.New()
	return &Log{logger: log}, nil
}

//LogInfo : log as info
func (l *Log) LogInfo(data interface{}) {
	l.logger.Info(data)
}
