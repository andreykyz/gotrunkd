// server.go
package main

import (
	"fmt"
	"net"
)

func server(connectInfo *ConnectInfo) {
	var len int
	var err error
	var mytrunkData TrunkData
	errorHandler := ErrorHandler{connectInfo.logger}
	connectInfo.logger.Info(fmt.Sprintf("Listening %d", connectInfo.port))
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", connectInfo.port))
	errorHandler.checkError(err)
	for {
		connectInfo.logger.Info("Waiting for clients...")
		conn, err := listener.Accept()
		connectInfo.logger.Info("Got some")
		go func(conn net.Conn) {
			b := make([]byte, 100)
			len, err = conn.Write([]byte(connectInfo.title))
			stage := ST_HOST
			for {
				len, err = conn.Read(b)
				errorHandler.checkError(err)
				connectInfo.logger.Info((string(b[0:len])))
				if err != nil {
					break
				}
				switch stage {
				case ST_HOST:
					if string(b[0:5]) == "HOST:" {
						var ok bool
						tmpString := string(b[5:len])
						connectInfo.logger.Debug(string(b[0:len]))
						_, ok = connectInfo.trunkData[tmpString]
						println(ok)
						if ok == false {
							connectInfo.logger.Err(fmt.Sprintf("Bad client: %s", tmpString))
							//	break
						}
						mytrunkData := connectInfo.trunkData[tmpString]
						println(mytrunkData.ipFrom)
						conn.Write([]byte("HOST OK"))
						stage = ST_CHAL
						continue
					}
					if string(b[0:5]) != "HOST:" {
						connectInfo.logger.Err("D_NOHOST")
						return
					}

				case ST_CHAL:
					if string(b[0:5]) == "CHAL:" {
						if mytrunkData.password != string(b[5:len]) {
							connectInfo.logger.Err("Password incorrect")
							break
						}
						conn.Write([]byte("CHAL OK"))
						stage = ST_PROT
						continue
					}
					if string(b[0:5]) != "CHAL:" {
						connectInfo.logger.Err("Password expected")
						break
					}
				case ST_PROT:
					if string(b[0:5]) == "PROT:" {
						connectInfo.protocol = string(b[5:7])
						stage = ST_TRUNK
						break
					}
				}
			}
			listener, err := net.Listen(connectInfo.protocol, "")
			errorHandler.checkError(err)
			conn.Write([]byte("PROT OK"))
			connectInfo.logger.Info("udp accept")
			mytrunkData.logicTunnel[0].conn, err = listener.Accept()
			errorHandler.checkError(err)
			connectInfo.logger.Info("udp accepted")

		}(conn)
	}

}
