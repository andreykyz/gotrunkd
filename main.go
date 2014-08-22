// gotrunkd project main.go
package main

import (
	//	"code.google.com/p/tuntap"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	connectInfo := new(ConnectInfo)
	flag.BoolVar(connectInfo.isServer, "server", false, "Run as server")
	flag.IntVar(connectInfo.port, "port", 5000, "Port number")
	flag.Parse()
	connectInfo.addr := net.ParseIP(flag.Arg(0))
	if connectInfo.addr == nil {
		fmt.Println("Invalid address")
		os.Exit(0)
	} else {
		fmt.Println("The address is ", connectInfo.addr.String())
	}
	println(connectInfo.isServer, connectInfo.port)

	if isServer {
		server(connectInfo)
	} else {
		client(connectInfo)
	}
}
