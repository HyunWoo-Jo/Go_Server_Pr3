package server_pr3

import (
	"Go_Server_Pr3/utills"
)

// 방 목록 요청
func RequestRoom(msgData MessageData) {

	roomDatas := rm.RoomList()
	// ../roomIpHash:userName:roomName
	requestMsg := "requestRoom"
	for _, roomData := range roomDatas {
		requestMsg += "/" + roomData.roomIpHash + ":" + roomData.userName + ":" + roomData.roomName
	}
	SendMessage(msgData.Conn, requestMsg)
}

// 방 생성
func CreateRoom(msgData MessageData, msg []string) {

	// msgs
	// ex) com:createRoom:3443:testUser:testRoom:1234
	// 0.command / 1.type / 2.port / 3.userName
	// 4.roomName / 5.password
	ip := utills.NetConnSplitIp(msgData.Conn) // ip port 분리
	cryptoKey := utills.CryptoSha256(ip)      // ip를 암호화

	roomData := RoomData{ // room Data 생성
		conn:       msgData.Conn,
		ip:         ip,
		roomIpHash: cryptoKey,
		port:       msg[2],
		userName:   msg[3],
		roomName:   msg[4],
		password:   msg[5],
	}

	// 방생성 RoomManager 등록
	rm.CreateRoom(cryptoKey, roomData)

	SendMessage(msgData.Conn, "sus create Room")
}

// 방 접속
func joinRoom(msgData MessageData, msg []string) {
	// msgs
	// 0.command / 1.type / 2.ipHash / 3.password
	roomData, isExist := rm.GetRoom(msg[2])
	if isExist {
		if roomData.password == msg[3] {
			roomDataMsg := "joinRoom/" + roomData.ip + ":" + roomData.port
			SendMessage(msgData.Conn, roomDataMsg)
		} else { // 비밀번호 일치하지 않음
			SendMessage(msgData.Conn, "Password doesn't match")
		}
	} else { // 방이 존재하지 않음
		SendMessage(msgData.Conn, "doesn't exist room")
	}

}

// 연결 취소
func Cancel(msgData MessageData) {
	msgData.Conn.Close()
}
