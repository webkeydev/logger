package logger

import (
	"fmt"
	"path"

	"github.com/sirupsen/logrus"
)

// ContextHook ...
type ContextHook struct{}

// Levels ...
func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire ...
func (hook ContextHook) Fire(entry *logrus.Entry) error {
	_, pkg := path.Split(path.Dir(entry.Caller.File))
	file := path.Base(entry.Caller.File)
	entry.Data["source"] = fmt.Sprintf("%s/%s:%v", pkg, file, entry.Caller.Line)
	return nil
}
