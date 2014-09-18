// connector
package main

import (
	"log"
)

type ConnectData struct {
	tunnelType   string
	chiperType   string
	compressType string
}

type ConnectInfo struct {
	title       string
	host        string
	port        int
	addr        string
	isServer    bool
	tun         string
	so_mark     int
	login       string
	pass        string
	logger      log.logger
	configPath  string
	connectData ConnectData
}

/* Authentication stage */
const (
	ST_INIT = iota
	ST_HOST
	ST_CHAL
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
)

func (connectInfo *ConnectInfo) loadFromNet(buf []byte, buf_len int) bool {

}
