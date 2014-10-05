// error.go
package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(1)
	}
}
