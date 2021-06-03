package connector

import (
	"sync"
	. "test/room"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Connector interface {
	AddConnection(conn Connection)
	SendMessageByRoom(roomID int64, data interface{})
}

type WSConnectorImpl struct {
	sync.RWMutex

	log         *logrus.Entry
	upgrader    *websocket.Upgrader
	rooms       map[int64]*RoomConnections
	roomService RoomService
}

func (c *WSConnectorImpl) AddConnection(conn Connection) {
	c.connect(conn)
}

func (c *WSConnectorImpl) SendMessageByRoom(roomID int64, data interface{}) {
	room := *c.rooms[roomID]
	if room == nil {
		c.log.Error("connector 32: room is nil")
		//TODO: some error here
		return
	}
	room.SendMessage(data)
}

func (c *WSConnectorImpl) connect(conn Connection) {
	err := conn.Connect()
	if err != nil {
		c.log.Error(err)
	}
	go c.listen(conn)
}

func (c *WSConnectorImpl) disconnect(conn Connection) {
	room := *c.rooms[conn.GetCurrentRoom()]
	room.RemoveConnection(conn)
}

func (c *WSConnectorImpl) listen(conn Connection) {
	for {
		select {
		case msg := <-conn.GetMessageChan():
			c.onEventMessage(conn, msg)
		case <-conn.GetDisconnectChan():
			c.disconnect(conn)
		}
	}
}

func NewWSConnector(log *logrus.Entry, roomService RoomService) Connector {
	return &WSConnectorImpl{
		log:         log,
		rooms:       map[int64]*RoomConnections{},
		roomService: roomService,
	}
}
