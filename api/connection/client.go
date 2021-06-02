package connection

import (
	"net/http"
	"test/user"

	"github.com/gorilla/websocket"
)

type Client struct {
	r             *http.Request
	conn          *websocket.Conn
	user          *user.User
	currentRoomID int64
	messages      chan []byte
	disconnect    chan struct{}
	close         chan struct{}
}
