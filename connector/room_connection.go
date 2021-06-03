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
	rc.Lock()
	defer rc.Unlock()
	rc.connections = append(rc.connections, conn)
}

func (rc *RoomConnectionsImpl) RemoveConnection(conn Connection) {
	rc.Lock()
	defer rc.Unlock()
	var newConn []Connection
	for _, roomConn := range rc.connections {
		if roomConn != conn {
			newConn = append(newConn, roomConn)
		}
	}
	rc.connections = newConn
}

func (rc *RoomConnectionsImpl) SendMessage(data interface{}) {
	rc.RLock()
	defer rc.RUnlock()
	for _, conn := range rc.connections {
		err := conn.SendMessage(data)
		if err != nil {
			//TODO: some error here
			rc.RemoveConnection(conn)
		}
	}
}

func NewRoomConnections(roomID int64) RoomConnections {
	return &RoomConnectionsImpl{
		roomID:          roomID,
		connections:     make([]Connection, 0),
		unSubscribeChan: make(chan struct{}),
	}
}
