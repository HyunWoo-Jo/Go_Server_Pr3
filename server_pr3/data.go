package server_pr3

import "net"

type ServerData struct {
	conn     net.Conn
	ip       string
	port     string
	userName string
	roomName string
	password string
}
type MessageData struct {
	Msg  string
	Conn net.Conn
}
