package server_pr3

import "net"

type serverData struct {
	conn     net.Conn
	ip       string
	port     string
	userName string
	roomName string
	password string
}
