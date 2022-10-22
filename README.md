# Logger

Package **webkeydev/logger** contains custom formatter for Logrus. This formatter
will print tag, line number of source code and fields in human readable format.

## Usage
### Example usage
```go
package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/webkeydev/logger"
)

func main() {
	mainLogger := log.NewLogger("main")
	log.SetTxtLogger()
	secondLogger := mainLogger.WithFields(logrus.Fields{"animal": "walrus", "language": "go"})
	secondLogger.Info("hello")
}
```

```
 2009-11-10 23:00:00 INFO [main] [language: go, animal: walrus] sandbox1049604770/prog.go:12: hello
 ```


### Using as formatter
```go
package main

import (
	log "github.com/sirupsen/logrus"
	formatter "github.com/webkeydev/logger"
)

func main() {
	formatter.SetTxtFormatterForLogger(log.StandardLogger())
	log.Printf("hello")
}
```
```
2022-10-23 00:16:12 INFO [main] main.go:10: hello
```
