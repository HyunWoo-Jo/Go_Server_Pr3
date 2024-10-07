package server_pr3

import (
	"fmt"
	"net"
)

var rm *RoomManager // room data를 관리 하는 Map

// 서버 Open / tcp 연결
func OpenServer(port string) {

	rm = NewRoomManager() // room manager 생성

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
		fmt.Println(conn.RemoteAddr().String())

		go OnReceiveMessage(conn)
	}
}
