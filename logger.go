package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = &logrus.Logger{
	Out:          os.Stdout,
	Level:        logrus.InfoLevel,
	Hooks:        make(logrus.LevelHooks),
	ReportCaller: true,
	Formatter: &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	},
}

// NewLogger creates a new context logger
func NewLogger(tag string) *logrus.Entry {
	return logger.WithField("tag", tag)
}

// Logger exports the main logger
func Logger() *logrus.Logger {
	return logger
}

// SetLoggerLevel sets the logging level of the main logger
func SetLoggerLevel(lvl string) error {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return err
	}
	logger.SetLevel(level)
	return nil
}

// SetJsonLogger sets the json logging format of the main logger
func SetJsonLogger() {
	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

// SetTxtLogger sets the text logging format of the main logger
func SetTxtLogger() {
	logger.Formatter = &MyFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LevelDesc:       []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG"},
	}
	logger.AddHook(&ContextHook{})
}
