package server_pr3_testing

import (
	"Go_Server_Pr3/utills"
	"net"
	"strconv"
	"time"
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
	SendMessage(conn, "com:createRoom")
	SendMessage(conn, "com:requestRoom")
}

// 수신
func OnReceiveMessage(conn net.Conn) {
	defer conn.Close() // 수신 에러시 종료
	for {
		data := make([]byte, 4096)
		message, err := conn.Read(data)
		if err != nil { // 에러
			utills.ColorPrintlnGreen("testing Read Err", err.Error())
			return
		}
		utills.ColorPrintlnGreen("testing Read: ", string(data[:message]), " Size: ", strconv.Itoa(len(string(data[:message]))))
	}
}

// 송신
func SendMessage(userConn net.Conn, message string) {

	utills.ColorPrintlnGreen("testing send :", message)
	_, err := userConn.Write([]byte(message))
	if err != nil {
		utills.ColorPrintlnGreen("testing send err ", err.Error())
	}
}
