// error.go
package main

import (
	"fmt"
	"log/syslog"
	"os"
)

type ErrorHandler struct {
	logger *syslog.Writer
}

func (errorHandler *ErrorHandler) checkError(err error) {
	if err != nil {
		errorHandler.logger.Err(fmt.Sprintf("Fatal error: %s", err.Error()))
		os.Exit(1)
	}
}
