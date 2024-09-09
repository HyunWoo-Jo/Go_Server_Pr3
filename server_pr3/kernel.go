package server_pr3

import (
	"Go_Server_Pr3/utills"
)

// 접근 유형 확인 후 분배
func OnKernel(msgData MessageData) {
	msgs := utills.Decoposit(msgData.Msg)
	switch msgs[1] {
	case "createRoom":
		CreateRoom(msgData, msgs)
	case "requestRoom":
		RequestRoom(msgData)
	case "joinRoom":
		joinRoom(msgData, msgs)
	case "cancel":
		Cancel(msgData)
	}

}
