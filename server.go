// server.go
package main

import (
	"log/syslog"
	"net"
	"strconv"
)

func server(connectInfo *ConnectInfo) {
	var len int
	var TERM int
	//	connectInfo.logger, err := syslog.NewLogger(syslog.LOG_WARNING|syslog.LOG_INFO|syslog.LOG_DEBUG, 0)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(connectInfo.port))
	checkError(err)
	buff := make([]byte, 1500)
	for TERM == 0 {
		conn, err := listener.Accept()
		go func(Listener conn) {
			b := make([]byte, 50)
			len, err = conn.write(connectInfo.title)
			stage = ST_HOST
			for {
				len, err = conn.Read(b)
				checkError(err)
				if err != nil {
					break
				}
				switch stage {
				case ST_HOST:
					if b[0:5] == "HOST:" {
						copy(connectInfo.host, b[5:])
						stage = ST_CHAL
						continue
					}
					if b[0:5] != "HOST:" {
						println("D_NOHOST")
						return
					}

				case ST_CHAL:
					if b[0:5] == "CHAL:" {
						copy(connectInfo.pass, b[5:])
						break
					}
					if b[0:5] != "CHAL:" {
						println("D_PWD")
						return
					}
				}
			}

		}(conn)
	}

}
