package connector

import (
	"encoding/json"
	"strconv"
)

type EventType int64

const (
	SubscribeEventType       EventType = 1 // Connect to room		Data: "1","2","3"
	UnSubscribeRoomEventType EventType = 2 // Disconnect from room  Data: ""
	NewMessageEventType      EventType = 3 // Message				Data: "Text Message"
)

type EventMessage struct {
	Type EventType `json:"type"`
	Data string    `json:"data"`
}

func (c *WSConnectorImpl) onEventMessage(conn Connection, msg []byte) {
	var message EventMessage
	err := json.Unmarshal(msg, &message)
	if err != nil {
		c.log.Error(err)
		//TODO: some error here
		return
	}
	switch message.Type {
	case SubscribeEventType:
		roomNumber, err := strconv.ParseInt(message.Data, 10, 64)
		if err != nil {
			c.log.Error(err)
			//TODO: some error here
			return
		}
		conn.SetCurrentRoom(roomNumber)
		if c.rooms[roomNumber] == nil {
			newRoom := NewRoomConnections(roomNumber)
			c.rooms[roomNumber] = &newRoom
		}
		room := *c.rooms[roomNumber]
		room.AddConnection(conn)
	case UnSubscribeRoomEventType:
		roomNumber, err := strconv.ParseInt(message.Data, 10, 64)
		if err != nil {
			c.log.Error(err)
			//TODO: some error here
			return
		}
		room := *c.rooms[roomNumber]
		room.RemoveConnection(conn)
	case NewMessageEventType:
		c.SendMessageByRoom(conn.GetCurrentRoom(), message)
	}
}
