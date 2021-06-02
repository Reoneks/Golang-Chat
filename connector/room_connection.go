package connector

import "sync"

type RoomConnections interface {
	AddConnection(conn Connection)
	RemoveConnection(conn Connection)
	SendMessage(data interface{})
}

type RoomConnectionsImpl struct {
	sync.RWMutex
	roomID          int64
	connections     []Connection
	unSubscribeChan chan struct{}
}

func (rc *RoomConnectionsImpl) AddConnection(conn Connection) {

}

func (rc *RoomConnectionsImpl) RemoveConnection(conn Connection) {

}

func (rc *RoomConnectionsImpl) SendMessage(data interface{}) {

}

func NewRoomConnections(roomID int64) RoomConnections {
	return &RoomConnectionsImpl{
		roomID:          roomID,
		connections:     make([]Connection, 0),
		unSubscribeChan: make(chan struct{}),
	}
}
