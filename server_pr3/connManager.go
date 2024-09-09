package server_pr3

import (
	"net"
	"sync"
)

// Conn 접속 데이터를 Map에 저장 관리
type ConnManager struct {
	mutex    sync.RWMutex
	connsMap map[string]net.Conn
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connsMap: make(map[string]net.Conn),
	}
}

func (connManager *ConnManager) AddConn(key string, conn net.Conn) {
	connManager.mutex.Lock()
	defer connManager.mutex.Unlock()
	connManager.connsMap[key] = conn
}

func (connManager *ConnManager) GetConn(key string) (net.Conn, bool) {
	connManager.mutex.Lock()
	defer connManager.mutex.Unlock()
	conn, exists := connManager.connsMap[key]
	return conn, exists
}

func (connManager *ConnManager) RemoveConn(key string) {
	connManager.mutex.Lock()
	defer connManager.mutex.Unlock()
	delete(connManager.connsMap, key)
}
