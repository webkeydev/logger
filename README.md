# Logger

Package webkeydev/logger contains custom formatter for Logrus. This formatter
will print tag, line number of source code and fileds in human readable formate.

## Usage
Example usage
```
package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/webkeydev/logger"
)

func main() {
	mainLogger := log.NewLogger("main")
	log.SetTxtLogger()
	mainLogger.WithFields(logrus.Fields{"animal": "walrus", "language": "go"})
}
```
