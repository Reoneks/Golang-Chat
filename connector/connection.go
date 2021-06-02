package connector

import (
	"github.com/gorilla/websocket"
	"net/http"
	. "test/user"
)

type Connection interface {
	GetMessageChan() chan []byte
	GetDisconnectChan() chan struct{}
	Connect() error
	Disconnect() error
	SendMessage(data interface{}) error
	ReadMessage() (messageType int, p []byte, err error)
	GetUser() *User
	GetCurrentRoom() int64
	SetCurrentRoom(roomID int64)
}

type WSConnection struct {
	r             *http.Request
	conn          *websocket.Conn
	user          *User
	currentRoomID int64

	isConnected    bool
	messageChan    chan []byte
	disconnectChan chan struct{} // on disconnect
	closeConnChan  chan struct{} // on close
}

func (c *WSConnection) GetMessageChan() chan []byte {
	return c.messageChan
}

func (c *WSConnection) GetDisconnectChan() chan struct{} {
	return c.disconnectChan
}

func (c *WSConnection) Connect() error {
	if c.isConnected {
		return nil
	}
	c.isConnected = true
	go c.connect()
	return nil
}

func connect() {
	for {
		select {
		case <-c.closeConnChan:
			c.isConnected = false
			break
		default:
			_, msgData, err := c.conn.ReadMessage()
			if err != nil {
				c.isConnected = false
				c.disconnectChan <- struct{}{}
				break
			}
			c.messageChan <- msgData
		}
	}

}

func (c *WSConnection) Disconnect() error {
	c.closeConnChan <- struct{}{}
	return c.conn.Close()
}

func (c *WSConnection) SendMessage(data interface{}) error {
	return c.conn.WriteJSON(data)
}

func (c *WSConnection) ReadMessage() (messageType int, p []byte, err error) {
	return c.conn.ReadMessage()
}

func (c *WSConnection) GetUser() *User {
	return c.user
}

func (c *WSConnection) GetCurrentRoom() int64 {
	return c.currentRoomID
}

func (c *WSConnection) SetCurrentRoom(roomID int64) {
	c.currentRoomID = roomID
}

func NewWSConnection(r *http.Request, conn *websocket.Conn, user *User) Connection {
	return &WSConnection{
		r:              r,
		conn:           conn,
		user:           user,
		currentRoomID:  0,
		isConnected:    false,
		messageChan:    make(chan []byte),
		disconnectChan: make(chan struct{}),
		closeConnChan:  make(chan struct{}),
	}
}
