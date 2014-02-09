# umisama/golog
golog is yet another logger implementation for golang.

## Features
 * 4 log level : DEBUG, INFO, WARN, CRITICAL
  * very small cost on non-output function call.
 * output format is free with text/template. can use these factors...
  * file name( full / short )
  * function name( full / short )
  * line number
  * time
  * message
 * time format is free

## Usage
[godoc is here](http://godoc.org/github.com/umisama/golog)

simple example:

```go
package main

import (
	"github.com/umisama/golog"
	"os"
)

var logger log.Logger

func OtherFuncName() {
	logger.Info("here is OtherFuncName()")
}

func main() {
	logger, _ = log.NewLogger(os.Stdout,
						log.TIME_FORMAT_SEC,		// Set time writting format.
						log.LOG_FORMAT_POWERFUL,	// Set log writting format.
						log.LogLevel_Debug)			// Set log level.

	logger.Debug("debug")
	logger.Info("Infomation")
	logger.Warn("Warning!")
	logger.Critical("Critical!")

	OtherFuncName()
}
```

```output
2014/2/9 12:41:59 main.go:20(main) : "debug"
2014/2/9 12:41:59 main.go:21(main) : "Infomation"
2014/2/9 12:41:59 main.go:22(main) : "Warning!"
2014/2/9 12:41:59 main.go:23(main) : "Critical!"
2014/2/9 12:41:59 main.go:11(OtherFuncName) : "here is OtherFuncName()"
```

in production, change  log level.

```go
	logger, _ = log.NewLogger(os.Stdout,
						log.TIME_FORMAT_SEC,		// Set time writting format.
						log.LOG_FORMAT_POWERFUL,	// Set log writting format.
						log.LogLevel_Critical)		// Set log level.
```

```output
2014/2/9 12:41:59 main.go:23(main) : "Critical!"
```

## install

```
go get github.com/umisama/golog
```

## Feedback
@umisama  
umisama@fe2o3.jp
