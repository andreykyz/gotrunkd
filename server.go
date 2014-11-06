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
			for stage != ST_TRUNK {
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
						if ok == false {
							connectInfo.logger.Err(fmt.Sprintf("Bad client: %s", tmpString))
							break
						}
						connectInfo.logger.Err(fmt.Sprintf("Good client: %s", tmpString))
						mytrunkData = connectInfo.trunkData[tmpString]
						conn.Write([]byte("HOST OK"))
						stage = ST_CHAL
						continue
					}
					if string(b[0:5]) != "HOST:" {
						connectInfo.logger.Err("D_NOHOST")
						return
					}

				case ST_CHAL:
					if string(b[0:5]) == "PASS:" {
						if mytrunkData.password != string(b[5:len]) {
							connectInfo.logger.Err(fmt.Sprintf("Password:%s incorrect", string(b[5:len])))
							break
						}
						conn.Write([]byte("PASS OK"))
						stage = ST_PROT
						continue
					}
					if string(b[0:5]) != "PASS:" {
						connectInfo.logger.Err("Password expected")
						break
					}
				case ST_PROT:
					if string(b[0:5]) == "PROT:" {
						connectInfo.protocol = string(b[5:8])
						stage = ST_TRUNK
						break
					}
				}
			}
			addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:")
			errorHandler.checkError(err)
			udpConn, err := net.ListenUDP(connectInfo.protocol, addr)
			errorHandler.checkError(err)
			connectInfo.logger.Info(fmt.Sprintf("Logic channel listening on %s", udpConn.LocalAddr().String()))
			conn.Write([]byte("PROT OK" + udpConn.LocalAddr().String()))
			_, rAddr, err := udpConn.ReadFromUDP(b)
			errorHandler.checkError(err)
			udpConn.WriteToUDP([]byte("ping from server"), rAddr)
			connectInfo.logger.Info(fmt.Sprintf("udp recv from %s", rAddr.String()))
			errorHandler.checkError(err)

		}(conn)
	}

}
