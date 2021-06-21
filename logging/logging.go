package logging

import (
	"io"
	logging "log"
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var logFileName = "logs/application.log"

var (
	log *logrus.Logger
)

func init() {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		logging.Fatalf("error opening file: %v", err)
	}

	log = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}

	log.SetReportCaller(true)

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.Info("--------------------------------------------------------")
	log.Info("Start logging")
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
