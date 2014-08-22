// connector
package main

type ConnectInfo struct {
	name     string
	port     int
	addr     string
	isServer bool
	tun      string
	so_mark  int
	login    string
	pass     string
}
