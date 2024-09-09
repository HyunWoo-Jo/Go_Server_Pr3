package server_pr3

// 방 목록 요청
func RequestRoom(msgData MessageData) {

}

// 방 생성
func CreateRoom(msgData MessageData) {

}

// 연결 취소
func Cancel(msgData MessageData) {
	msgData.Conn.Close()
}
