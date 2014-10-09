// client
package main

import (
	//	"encoding/binary"
	"fmt"
	"log/syslog"
	"net"
	_ "os"
	"strconv"
	_ "unsafe"
)

func client(connectInfo *ConnectInfo, myName string) {
	var len int
	var stage int = ST_INIT
	var err error
	myTrunkData := connectInfo.trunkData[myName]
	myTrunkData.logger, err = syslog.New(syslog.LOG_WARNING|syslog.LOG_INFO|syslog.LOG_DEBUG, fmt.Sprintf("gotrunk %s", myName))
	myTrunkData.logger.Info(fmt.Sprintf("Connecting to %s:%d", connectInfo.addr, connectInfo.port))
	errorHandler := ErrorHandler{myTrunkData.logger}
	conn, err := net.Dial("tcp", connectInfo.addr+":"+strconv.Itoa(connectInfo.port))
	errorHandler.checkError(err)
	b := make([]byte, 100)
	for {
		len, err = conn.Read(b)
		errorHandler.checkError(err)
		myTrunkData.logger.Info((string(b[0:len])))
		if err != nil {
			myTrunkData.logger.Err("D_NOREAD")
			break
		}
		if len == 0 {
			myTrunkData.logger.Err("D_NOREAD")
			break
		}
		switch stage {
		case ST_INIT:
			if string(b[0:7]) != "VTRUNKD" {
				myTrunkData.logger.Err("D_NOREAD")
				return
			}
			len, err = conn.Write([]byte("HOST:" + myName))
			stage = ST_HOST
			continue
		case ST_HOST:
			if string(b[0:7]) != "HOST OK" {
				myTrunkData.logger.Err("D_NOHOST")
				return
			}
			len, err = conn.Write([]byte("CHAL:" + myTrunkData.password))
			stage = ST_CHAL

		case ST_CHAL:
			if string(b[0:7]) != "CHAL OK" {
				myTrunkData.logger.Err("D_PWD")
				return
			}
			len, err = conn.Write([]byte("PROT:udp"))
			stage = ST_PROT
		case ST_PROT:
			if string(b[0:7]) != "PROT OK" {
				myTrunkData.logger.Err("D_PROT")
				return
			}
			//len, err = conn.Write([]byte("GET PORTS"))
			break
			//			connectInfo.trunkData.port = int(binary.BigEndian.Uint16(b[7:9]))

		}
	}

}
