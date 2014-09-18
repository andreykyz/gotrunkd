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
	var stage int = ST_INIT
	var reason = D_NOREAD
	conn, err := net.Dial("tcp", connectInfo.addr+":"+strconv.Itoa(connectInfo.port))
	checkError(err)
	b := make([]byte, 50)
	for {
		len, err = conn.Read(b)
		checkError(err)
		if err != nil {
			break
		}

		switch stage {
		case ST_INIT:
			if b[0:7] != "VTRUNKD" {
				println("D_NOREAD")
				return
			}
			len, err = conn.Write("HOST:" + connectInfo.host + "\n")
			stage = ST_HOST
			continue
		case ST_HOST:
			if b[0:2] != "OK" {
				println("D_NOHOST")
				return
			}
			len, err = conn.Write("CHAL:" + connectInfo.pass + "\n")
			stage = ST_CHAL

		case ST_CHAL:
			if b[0:2] != "OK" {
				println("D_PWD")
				return
			}
		}
	}

}
