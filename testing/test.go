package server_pr3_testing

import (
	"Go_Server_Pr3/utills"
	"net"
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
	defer conn.Close()
	utills.ColorPrintlnGreen("conneting")
}
