package utility

import (
	"net/http"

	"github.com/mandocaesar/go-skeleton/config"
	"github.com/sebest/logrusly"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus"
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
		"service": cfg.Server.Name,
		"version": cfg.Server.Version,
	})

	hook := logrusly.NewLogglyHook(config.Token, config.Host, logrus.InfoLevel, config.Tags...)
	log.Hooks.Add(hook)
	log.Hooks.Add(&apmlogrus.Hook{})

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

func handleRequest(w http.ResponseWriter, req *http.Request) {
	traceContextFields := apmlogrus.TraceContext(req.Context())
	logrus.WithFields(traceContextFields).Debug("handling request")
}
