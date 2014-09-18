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
	flag.StringVar(&connectInfo.configPath, "config", "/etc/gotrunkd.default.config", "Path to config file")

	flag.IntVar(&connectInfo.port, "port", 5000, "Port number")

	flag.Parse()
	connectInfo.addr = net.ParseIP(flag.Arg(0)).String()

	println(connectInfo.isServer, connectInfo.addr, ":", connectInfo.port)
	// config file parse
	c, err := conf.ReadConfigFile(connectInfo.configPath)
	if connectInfo.isServer {
		connectInfo.title = "VTRUNKD server version 0.1go"
		server(connectInfo)
	} else {
		connectInfo.title = "VTRUNKD client version 0.1go"
		go client(connectInfo)
	}
}
