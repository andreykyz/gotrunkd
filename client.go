// client
package main

import (
	"net"
	_ "os"
	"strconv"
	_ "unsafe"
)

func client(connectInfo *ConnectInfo) {
	var len int
	conn, err := net.Dial("tcp", connectInfo.addr+":"+strconv.Itoa(connectInfo.port))
	checkError(err)
	b := make([]byte, 1500)
	copy(b[0:], connectInfo.login)
	len, err = conn.Write(b[0:len])
	len, err = conn.Read(b)
	checkError(err)
}
