package server_pr3_testing

import (
	"Go_Server_Pr3/utills"
	"bufio"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	requestListener = make(chan string)
	joinListener    = make(chan string)
)

func OnTest_Server(port string) {
	time.Sleep(3000) // 3초 지연
	utills.ColorPrintlnGreen("start Test")
	conn, err := net.Dial("tcp", "127.0.0.1"+port)
	if err != nil {
		utills.ColorPrintlnGreen("test conecting err:", err.Error())
		return
	}
	utills.ColorPrintlnGreen("conneting")
	go OnReceiveMessage(conn) // 수신 대기
	time.Sleep(1000)          // 1초 지연
	SendMessage(conn, "com:createRoom:3443:testUser:testRoom0:1234")
	time.Sleep(1000)
	SendMessage(conn, "com:requestRoom")

	utills.ColorPrintlnGreen("waiting..")
	requestRoom := <-requestListener
	msgs := strings.Split(requestRoom, "/")
	roomDatas := strings.Split(msgs[1], ":")

	hashIp := roomDatas[0]
	cmd := "com:joinRoom:" + hashIp + ":"
	SendMessage(conn, cmd)
	time.Sleep(1000)
	cmd += "12"
	SendMessage(conn, cmd)
	time.Sleep(1000)
	cmd += "34"
	SendMessage(conn, cmd)
}

// 수신
func OnReceiveMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil { // 에러
			utills.ColorPrintlnGreen("testing Read Err", err.Error())
			return
		}
		utills.TrimNewline(&message)
		utills.ColorPrintlnGreen("testing Read: ", message, " Size: ", strconv.Itoa(len(message)))

		switch strings.Split(message, "/")[0] {
		case "requestRoom":
			requestListener <- message
		case "joinRoom":
			joinListener <- message
		}
	}
}

// 송신
func SendMessage(userConn net.Conn, message string) {

	utills.ColorPrintlnGreen("testing send: ", message)
	_, err := userConn.Write([]byte(message + "\n"))
	if err != nil {
		utills.ColorPrintlnGreen("testing send err ", err.Error())
	}
}
