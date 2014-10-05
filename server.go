// server.go
package main

import (
	"net"
	"strconv"
)

func server(connectInfo *ConnectInfo) {
	var len int
	var err error
	var mytrunkData TrunkData
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(connectInfo.port))
	checkError(err)
	for {
		conn, err := listener.Accept()
		go func(conn net.Conn) {
			b := make([]byte, 50)
			len, err = conn.Write([]byte(connectInfo.title))
			stage := ST_HOST
			for {
				len, err = conn.Read(b)
				checkError(err)
				println(string(b[0:10]))
				if err != nil {
					break
				}
				switch stage {
				case ST_HOST:
					if string(b[0:5]) == "HOST:" {
						tmpString := string(b[5:])
						_, ok := connectInfo.trunkData[tmpString]
						if ok == false {
							println("Bad client")
							break
						}
						mytrunkData := connectInfo.trunkData[tmpString]
						mytrunkData.name = tmpString
						conn.Write([]byte("HOST OK"))
						stage = ST_CHAL
						continue
					}
					if string(b[0:5]) != "HOST:" {
						println("D_NOHOST")
						return
					}

				case ST_CHAL:
					if string(b[0:5]) == "CHAL:" {
						mytrunkData.password = string(b[5:])
						conn.Write([]byte("CHAL OK"))
						stage = ST_PROT
						continue
					}
					if string(b[0:5]) != "CHAL:" {
						println("D_PWD")
						break
					}
				case ST_PROT:
					if string(b[0:5]) == "PROT:" {
						connectInfo.protocol = string(b[5:7])
						conn.Write([]byte("PROT OK"))
						stage = ST_TRUNK
						continue
					}

				case ST_TRUNK:

				}
			}

		}(conn)
	}

}
