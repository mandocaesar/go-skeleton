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
	config *config.Configuration
}

//NewLogger : instantiate new logger
func NewLogger(cfg *config.Configuration) (*Log, error) {
	_log := logrus.New()
	_config := cfg

	_log.WithFields(logrus.Fields{
		"service": cfg.Server.Name,
		"version": cfg.Server.Version,
	})

	hook := logrusly.NewLogglyHook(_config.Loggly.Token, _config.Loggly.Host, logrus.InfoLevel, _config.Loggly.Tags...)
	_log.Hooks.Add(hook)
	_log.Hooks.Add(&apmlogrus.Hook{})

	return &Log{logger: _log, config: _config}, nil
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
