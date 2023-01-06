package logger

import (
	"github.com/sirupsen/logrus"
)

// NewLogger creates a new context logger
func NewLogger(tag string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{"tag": tag})
}

// SetLoggerLevel sets the logging level of the main logger
func SetLoggerLevel(lvl string) error {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return err
	}
	logrus.StandardLogger().SetLevel(level)
	return nil
}

// SetJsonLogger sets the json logging format of the std logger
func SetJsonLogger() {
	std := logrus.StandardLogger()
	std.ReportCaller = true
	std.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

// SetTxtLogger sets the text logging format of the std logger
func SetTxtLogger() {
	std := logrus.StandardLogger()
	std.Formatter = NewMyFormatter()
	std.ReportCaller = true
	std.AddHook(&ContextHook{})
}

// SetTxtFormatterForLogger prepare txt logger for the given logger
func SetTxtFormatterForLogger(logger *logrus.Logger) {
	logger.Formatter = NewMyFormatter()
	logger.ReportCaller = true
	logger.AddHook(&ContextHook{})
}
