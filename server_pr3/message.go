package server_pr3

import (
	"fmt"
	"net"
)

var (
	readListener = make(chan MessageData)
)

// 송신
func SendMessage(userConn net.Conn, message string) {

	fmt.Println("send :", message)
	_, err := userConn.Write([]byte(message))
	if err != nil {
		fmt.Println("send err ", err)
	}
}

// 수신
func OnReceiveMessage(conn net.Conn) {
	for {
		data := make([]byte, 4096)
		message, err := conn.Read(data)
		if err != nil { // 에러
			fmt.Println("Read Err", err)
			cancel := MessageData{"err:cancel", conn}
			readListener <- cancel
			return
		}
		fmt.Printf("Read : %s Size: %d\n", string(data[:message]), len(string(data[:message])))
		readListener <- MessageData{string(data[:message]), conn}
	}
}
