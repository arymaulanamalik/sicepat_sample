package logger

import (
	"context"
	"strings"

	"gitlab.sicepat.tech/platform/golib/log"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Entry

//Configure configure log
func Configure() {

	format := "json"
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	lvl := strings.ToLower("debug")
	switch lvl {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	Log = logrus.WithFields(logrus.Fields{
		"App": "authorization-svc",
	})
}

func setDefault() *logrus.Entry {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	return logrus.WithField("app", "authorization-svc")
}

//GetLogger get logger instance
func GetLogger(pkg, funcName string) *logrus.Entry {
	if Log == nil {
		setDefault()
	}
	return log.WithFields(log.Fields{
		"function":   funcName,
		"package":    pkg,
		"request-id": uuid.New().String(),
	})
}

//GetLoggerContext get logger with context
func GetLoggerContext(ctx context.Context, pkg, funcName string) *logrus.Entry {
	if Log == nil {
		setDefault()
	}

	return log.WithContext(ctx).WithFields(logrus.Fields{
		"package":    pkg,
		"function":   funcName,
		"request-id": uuid.New().String(),
	})
}
