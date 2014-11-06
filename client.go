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
	for stage != ST_TRUNK {
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
			myTrunkData.logger.Info(string(b[0:7]))
			len, err = conn.Write([]byte("PASS:" + myTrunkData.password))
			stage = ST_CHAL

		case ST_CHAL:
			if string(b[0:7]) != "PASS OK" {
				myTrunkData.logger.Err("D_PWD")
				return
			}
			myTrunkData.logger.Info(string(b[0:7]))
			len, err = conn.Write([]byte("PROT:udp"))
			stage = ST_PROT
		case ST_PROT:
			if string(b[0:7]) != "PROT OK" {
				myTrunkData.logger.Err("D_PROT")
				return
			}
			myTrunkData.logger.Info(fmt.Sprintf("Dial on %s", string(b[7:len])))
			raddr_str := string(b[7:len])
			raddr, err := net.ResolveUDPAddr("udp", raddr_str)
			errorHandler.checkError(err)
			myTrunkData.logicTunnel.conn, err = net.DialUDP("udp", nil, raddr)
			errorHandler.checkError(err)
			myTrunkData.logger.Info("Dial ok")
			len, err = myTrunkData.logicTunnel.conn.Write([]byte(fmt.Sprintf("Im %s", myTrunkData.name)))
			errorHandler.checkError(err)
			len, err = myTrunkData.logicTunnel.conn.Read(b)
			errorHandler.checkError(err)
			myTrunkData.logger.Info(string(b[0:len]))

			break
			//			connectInfo.trunkData.port = int(binary.BigEndian.Uint16(b[7:9]))

		}
	}

}
