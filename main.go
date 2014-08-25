// gotrunkd project main.go
package main

import (
	//	"code.google.com/p/tuntap"
	"code.google.com/p/goconf/conf"
	"flag"
	_ "fmt"
	"net"
	_ "os"
)

func main() {
	connectInfo := new(ConnectInfo)
	// comand line parse
	flag.BoolVar(&connectInfo.isServer, "server", false, "Run as server")
	flag.IntVar(&connectInfo.port, "port", 5000, "Port number")
	flag.Parse()
	connectInfo.addr = net.ParseIP(flag.Arg(0)).String()

	println(connectInfo.isServer, connectInfo.addr, ":", connectInfo.port)
	// config file parse
	c, err := conf.ReadConfigFile("something.config")
	//	c.
	if connectInfo.isServer {
		server(connectInfo)
	} else {
		client(connectInfo)
	}
}
