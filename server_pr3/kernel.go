package server_pr3

import (
	"Go_Server_Pr3/utills"
)

// 접근 유형 확인 후 분배
func OnKernel(msgData MessageData) {
	msg := utills.Decoposit(msgData.Msg)
	switch msg[1] {
	case "createRoom":
		CreateRoom(msgData)
	case "requestRoom":
		RequestRoom(msgData)
	case "cancel":
		Cancel(msgData)
	}

}
