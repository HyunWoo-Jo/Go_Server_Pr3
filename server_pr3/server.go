package server_pr3

import (
	"fmt"
	"net"
)

// 서버 Open / tcp 연결
func OpenServer(port string) {
	lnsten, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Server err :", err)
		return
	}
	defer lnsten.Close()

	go OnAccept(lnsten)

	// Server Off 처리
	s := ""
	for {
		fmt.Println("Server Off command \"server -off\"")
		fmt.Scanln(&s)
		if s == "server -off" {
			break
		}
	}
	fmt.Println("close")
}

// 연결 수락
func OnAccept(ln net.Listener) {
	fmt.Println("On Accept")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept err :", err)
			continue
		}
		go ReceiveMessage(conn)
	}
}

// 연결 취소
func Cancel(conn net.Conn) {
	conn.Close()

}
