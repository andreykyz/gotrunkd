// connector
package main

import (
	"log/syslog"
)

type TrunkData struct {
	tunnelType   string
	chiperType   string
	password     string
	compressType string
	name         string
	conn         []LogicTunnel
	connAmount   int
	tun          string
	ipFrom       string
	ipTo         string
}
type LogicTunnel struct {
	port int
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

func (connectInfo *ConnectInfo) loadFromFile() bool {

	return true
}
func (connectInfo *ConnectInfo) loadFromNet(buf []byte, buf_len int) bool {
	return true
}
