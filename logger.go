package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var std *logrus.Logger

func init() {
	std = logrus.StandardLogger()
	std.Out = os.Stdout
	std.ReportCaller = true
	std.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

// NewLogger creates a new context logger
func NewLogger(tag string) *logrus.Entry {
	return std.WithField("tag", tag)
}

// SetLoggerLevel sets the logging level of the main logger
func SetLoggerLevel(lvl string) error {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return err
	}
	std.SetLevel(level)
	return nil
}

// SetJsonLogger sets the json logging format of the main logger
func SetJsonLogger() {
	std.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

// SetTxtLogger sets the text logging format of the main logger
func SetTxtLogger() {
	std.Formatter = &MyFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LevelDesc:       []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG"},
	}
	std.AddHook(&ContextHook{})
}
