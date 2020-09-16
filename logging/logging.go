package logging

import (
	"io"
	logging "log"
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	f, err := os.OpenFile("logs/application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		logging.Fatalf("error opening file: %v", err)
	}

	log = logrus.New()

	//log.Formatter = &logrus.JSONFormatter{}

	log.SetReportCaller(true)

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
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
