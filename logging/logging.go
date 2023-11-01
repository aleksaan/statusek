package logging

import (
	"math/rand"
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var RLogger *logrus.Entry

// var (
// 	log *logrus.Logger
// )

var log = logrus.New()

func init() {
	log.SetFormatter(new(prefixed.TextFormatter))
	log.SetOutput(os.Stdout)

	log.Info("--------------------------------------------------------")
	log.Info("Start logging")
}

func CreateRequestLogger(caller string, endpoint string) {
	RLogger = log.WithFields(logrus.Fields{"caller": caller, "endpoint": endpoint, "session_id": rand.Intn(1000000000)})
}

// Info ...
func Info(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Warn ...
func Warn(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Error ...
func Error(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

var (
	// ConfigError ...
	ConfigError = "%v type=config.error"

	// HTTPError ...
	HTTPError = "%v type=http.error"

	// HTTPWarn ...
	HTTPWarn = "%v type=http.warn"

	// HTTPInfo ...
	HTTPInfo = "%v type=http.info"
)
