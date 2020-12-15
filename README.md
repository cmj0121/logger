# Logger #
![Action](https://github.com/cmj0121/logger/workflows/Go%20test/badge.svg)

The simple Go-based log sub-system

## Example ##
The following is the sample code to use the logger in your project, you can change the log level by the
`logger.SetLevel` with pass the 

```go
package main

import (
	"github.com/cmj0121/logger"
)

func main() {
	log := logger.New("example")
	log.SetLevel("info")

	log.Crit("show the CRIT log")
	log.Warn("show the WARN log")
	log.Info("show the INFO log")
	log.Debug("show the DEBUG log")
	log.Verbose("show the VERBOSE log")
}
```
