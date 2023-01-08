package utils

import (
	"fmt"
	"github.com/jurelou/forensibus/utils/zap_logger"
)

var Log Logger

// Logger represent common interface for logging function
type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

func init() {
	var err error
	Log, err = zap_logger.GetLogger()
	if err != nil {
		panic(fmt.Sprintf("Could not create a logger : %s", err))
	}
}
