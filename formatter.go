package logger

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type MyFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var fields string
	keys := make([]string, 0, len(entry.Data))
	for k, v := range entry.Data {
		if k == "tag" {
			continue
		}

		if k == "source" {
			continue
		}
		keys = append(keys, fmt.Sprintf("%s: %v", k, v))
	}

	if len(keys) > 0 {
		fields = fmt.Sprintf("[%s] ", strings.Join(keys, ", "))
	}

	if _, ok := entry.Data["tag"]; !ok {
		entry.Data["tag"] = ""
	}

	return []byte(fmt.Sprintf("%s %s [%s] %s%s: %s\n", entry.Time.Format(f.TimestampFormat), f.LevelDesc[entry.Level], entry.Data["tag"], fields, entry.Data["source"], entry.Message)), nil
}
