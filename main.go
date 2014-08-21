// gotrunkd project main.go
package main

import (
	//	"code.google.com/p/tuntap"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	var isServer bool
	var port int
	flag.BoolVar(&isServer, "server", false, "Run as server")
	flag.IntVar(&port, "port", 5000, "Port number")
	flag.Parse()
	addr := net.ParseIP(flag.Arg(0))
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(0)
	} else {
		fmt.Println("The address is ", addr.String())
	}
	println(isServer, port)

	if isServer {

	} else {
		go client()
		conn, err := net.Dial("tcp", addr+":"+strconv.Itoa(port))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
