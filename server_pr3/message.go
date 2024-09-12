package server_pr3

import (
	"Go_Server_Pr3/utills"
	"bufio"
	"fmt"
	"net"
)

type MessageData struct {
	Msg  string
	Conn net.Conn
}

// 송신
func SendMessage(userConn net.Conn, message string) {

	fmt.Println("send :", message)
	_, err := userConn.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("send err ", err)
	}
}

// 수신
func OnReceiveMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil { // 에러
			fmt.Println("Read Err", err)
			cancel := MessageData{"err:cancel", conn}
			go OnKernel(cancel)
			return
		}
		utills.TrimNewline(&message)
		fmt.Printf("Read : %s Size: %d\n", message, len(message))
		go OnKernel(MessageData{message, conn})
	}
}
