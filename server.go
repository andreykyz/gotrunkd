// server.go
package main

import (
	"log/syslog"
	"net"
	"strconv"
)

func server(connectInfo ConnectInfo) {
	var len int
	list, err := net.Listen("tcp", ":"+strconv.Itoa(connectInfo.port))
	checkError(err)
	b := make([]byte, 1500)
	conn, err := listener.Accept()
	len, err = conn.Read(b)
	checkError(err)
	logger, err := NewLogger(LOG_WARNING | LOG_INFO | LOG_DEBUG,0)
}
