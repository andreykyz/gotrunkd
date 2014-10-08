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
		//		errorHandler.logger.
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(1)
	}
}
