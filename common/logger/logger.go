package logger

import (
	"go.uber.org/zap"
	"sync"
)

var (
	log  *zap.SugaredLogger
	once sync.Once
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Sync() error
}

// New returns logger interface.
// service name - will be added as value for log string
// env - dev/test value will use `Debug` level, default level is `Info`
func New(serviceName string, env string) Logger {
	once.Do(func() {
		var logger *zap.Logger
		if env == "dev" || env == "test" {
			logger, _ = zap.NewDevelopment()
		} else {
			logger, _ = zap.NewProduction()
		}

		log = logger.Sugar().Named(serviceName)
	})
	return log
}

//TODO: add service to store logs for services
