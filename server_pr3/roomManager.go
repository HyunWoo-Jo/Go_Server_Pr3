package server_pr3

import (
	"net"
	"sync"
)

type RoomData struct {
	conn       net.Conn
	ip         string
	port       string
	roomIpHash string
	userName   string
	roomName   string
	password   string
}

type RoomManager struct {
	mutex    sync.Mutex
	roomsMap map[string]RoomData // key room Number
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		roomsMap: make(map[string]RoomData),
	}
}

// 방 생성
func (roomManager *RoomManager) CreateRoom(key string, roomData RoomData) {
	roomManager.mutex.Lock()
	defer roomManager.mutex.Unlock()
	roomManager.roomsMap[key] = roomData
}

func (roomManager *RoomManager) GetRoom(key string) (RoomData, bool) {
	roomManager.mutex.Lock()
	defer roomManager.mutex.Unlock()
	roomData, isExist := roomManager.roomsMap[key]
	return roomData, isExist
}

// 방 제거
func (roomManager *RoomManager) RemoveRoom(key string) {
	roomManager.mutex.Lock()
	defer roomManager.mutex.Unlock()
	delete(roomManager.roomsMap, key)
}

// 방 목록을 반환
func (roomManager *RoomManager) RoomList() []RoomData {
	roomDatas := make([]RoomData, len(roomManager.roomsMap))
	i := 0
	for _, roomData := range roomManager.roomsMap {
		roomDatas[i] = roomData
		i++
	}
	return roomDatas
}
