package server_pr3

import (
	"Go_Server_Pr3/utills"
	"strconv"
)

// 방 목록 요청
func RequestRoom(msgData MessageData) {

	roomDatas := rm.RoomList()
	// ..:roomIpHash/roomName/userName/isPublic
	requestMsg := "data:roomList:"

	for _, roomData := range roomDatas {
		isPassword := roomData.password != "" // public ture, private false
		requestMsg += roomData.roomIpHash + "/" + roomData.roomName + "/" + roomData.userName + "/" + strconv.FormatBool(isPassword)
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
	SendMessage(msgData.Conn, "msg:sus create Room")
}

// 방 접속
func joinRoom(msgData MessageData, msg []string) {
	// msgs
	// 0.command / 1.type / 2.ipHash / 3.password
	roomData, isExist := rm.GetRoom(msg[2])
	if isExist {
		if roomData.password == msg[3] {
			roomDataMsg := "data:joinRoom:" + roomData.ip + "/" + roomData.port
			SendMessage(msgData.Conn, roomDataMsg)
		} else { // 비밀번호 일치하지 않음
			SendMessage(msgData.Conn, "msg:Password doesn't match")
		}
	} else { // 방이 존재하지 않음
		SendMessage(msgData.Conn, "msg:doesn't exist room")
	}

}

// 연결 취소
func Cancel(msgData MessageData) {
	ip := utills.NetConnSplitIp(msgData.Conn) // ip port 분리
	cryptoKey := utills.CryptoSha256(ip)      // ip 암호화
	_, isExist := rm.GetRoom(cryptoKey)
	if isExist {
		rm.RemoveRoom(cryptoKey)
	}
	msgData.Conn.Close()
}
