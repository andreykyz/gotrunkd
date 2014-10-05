// client
package main

import (
	//	"encoding/binary"
	"net"
	_ "os"
	"strconv"
	_ "unsafe"
)

func client(connectInfo *ConnectInfo, myName string) {
	var len int
	var stage int = ST_INIT
	myTrunkData := connectInfo.trunkData[myName]
	conn, err := net.Dial("tcp", connectInfo.addr+":"+strconv.Itoa(connectInfo.port))
	checkError(err)
	b := make([]byte, 50)
	for {
		len, err = conn.Read(b)
		checkError(err)
		println(string(b[0:10]))
		if err != nil {
			println("D_NOREAD")
			break
		}
		if len == 0 {
			println("D_NOREAD")
			break
		}
		switch stage {
		case ST_INIT:
			if string(b[0:7]) != "VTRUNKD" {
				println("D_NOREAD")
				return
			}
			len, err = conn.Write([]byte("HOST:" + myName + "\n"))
			stage = ST_HOST
			continue
		case ST_HOST:
			if string(b[0:7]) != "HOST OK" {
				println("D_NOHOST")
				return
			}
			len, err = conn.Write([]byte("CHAL:" + myTrunkData.password + "\n"))
			stage = ST_CHAL

		case ST_CHAL:
			if string(b[0:7]) != "CHAL OK" {
				println("D_PWD")
				return
			}
			len, err = conn.Write([]byte("PROT:udp" + "\n"))
			stage = ST_PROT
		case ST_PROT:
			if string(b[0:7]) != "PROT OK" {
				println("D_PROT")
				return
			}
			len, err = conn.Write([]byte("GET PORTS:" + "\n"))
			stage = ST_TRUNK
		case ST_TRUNK:

			//			connectInfo.trunkData.port = int(binary.BigEndian.Uint16(b[7:9]))

		}
	}

}
