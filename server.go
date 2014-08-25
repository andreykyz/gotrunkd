// server.go
package main

import (
	"log/syslog"
	"net"
	"strconv"
)

func server(connectInfo *ConnectInfo) {
	var len int
	logger, err := syslog.NewLogger(syslog.LOG_WARNING|syslog.LOG_INFO|syslog.LOG_DEBUG, 0)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(connectInfo.port))
	checkError(err)
	b := make([]byte, 1500)
	conn, err := listener.Accept()
	len, err = conn.Read(b)
	checkError(err)

}
