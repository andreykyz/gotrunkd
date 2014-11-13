// connector
package main

import (
	"code.google.com/p/tuntap"
	"encoding/binary"
	"fmt"
	"log/syslog"
	"net"
	"syscall"
)

type TrunkData struct {
	tunnelType   string
	chiperType   string
	password     string
	compressType string
	name         string
	logicTunnel  LogicTunnel
	connAmount   int
	tun          string
	ipFrom       string
	ipTo         string
	logger       *syslog.Writer
	dev          *tuntap.Interface
}
type LogicTunnel struct {
	raddr   string
	conn    *net.UDPConn
	fd      uintptr
	port    int
	tcpConn *net.TCPConn
}

type ConnectInfo struct {
	title      string
	routineNum int
	host       string
	port       int
	protocol   string
	addr       string
	isServer   bool
	tun        string
	so_mark    int
	password   string
	logger     *syslog.Writer
	configPath string
	trunkData  map[string]TrunkData
}

/* Authentication stage */
const (
	ST_INIT = iota
	ST_HOST
	ST_CHAL
	ST_PROT
	ST_TRUNK
)

/* Authentication errors */
const (
	D_NOSHAKE1 = iota
	D_NOSHAKE2
	D_ST_CHAL
	D_CHAL
	D_NOHOST
	D_NOMULT
	D_GREET
	D_PWD
	D_NOREAD
	D_OTHER
	D_PROT
)
const (
	FRAME_DATA             = 0x00
	FRAME_ECHO_REQ         = 0x01
	FRAME_ECHO_REP         = 0x02
	FRAME_CONN_CLOSE       = 0x03
	FRAME_CHANNEL_INFO     = 0x04
	FRAME_TIME_LAG         = 0x05
	FRAME_LAST_WRITTEN_SEQ = 0x06
)
const (
	TERM_FATAL = iota
	TERM_NONFATAL
	TERM_NON
)

func (connectInfo *ConnectInfo) loadFromFile() bool {

	return true
}
func (connectInfo *ConnectInfo) loadFromNet(buf []byte, buf_len int) bool {
	return true
}

// stolen from https://github.com/pebbe/zmq2
func FD_SET(p *syscall.FdSet, i uintptr) {
	p.Bits[i/64] |= 1 << uint(i) % 64
}
func FD_ISSET(p *syscall.FdSet, i uintptr) bool {
	return (p.Bits[i/64] & (1 << uint(i) % 64)) != 0
}
func FD_ZERO(p *syscall.FdSet) {
	for i := range p.Bits {
		p.Bits[i] = 0
	}
}

type LocalData struct {
}

func (connectInfo *ConnectInfo) trunk(trunkName string) {
	var err error
	running := 1
	//	var trunkData *TrunkData
	trunkData := connectInfo.trunkData[trunkName]
	logger := connectInfo.logger
	errorHandler := ErrorHandler{connectInfo.logger}
	trunkData.dev, err = tuntap.Open(connectInfo.tun, tuntap.DevTun)
	errorHandler.checkError(err)
	tunFd := trunkData.dev.File().Fd()

	var readFDSet syscall.FdSet
	//	var writeFDSet syscall.FdSet
	maxFD := int(tunFd)
	file, err := trunkData.logicTunnel.conn.File()
	errorHandler.checkError(err)
	trunkData.logicTunnel.fd = file.Fd()
	if maxFD < int(trunkData.logicTunnel.fd) {
		maxFD = int(trunkData.logicTunnel.fd)
	}
	var selectTime syscall.Timeval
	selectTime.Sec = 0
	selectTime.Usec = 10000

	data := make([]byte, 1500)

	for running == 1 {

		FD_ZERO(&readFDSet)
		FD_SET(&readFDSet, tunFd)
		FD_SET(&readFDSet, trunkData.logicTunnel.fd)
		selectRet, err := syscall.Select(maxFD+1, &readFDSet, nil, nil, &selectTime)
		errorHandler.checkError(err)
		if selectRet == 0 {
			trunkData.logger.Info("Idle...")
			continue
		}
		if FD_ISSET(&readFDSet, trunkData.logicTunnel.fd) {
			len, err := trunkData.logicTunnel.conn.Read(data)
			errorHandler.checkError(err)
			if len >= 0 && len < 2 {
				logger.Err(fmt.Sprintf("Net read return %d", len))
				running = 0
				break
			}
			packetFlag := binary.BigEndian.Uint16(data[0:2])
			switch packetFlag {
			case FRAME_DATA:

			}

		}

	}

}
