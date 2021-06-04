package connector

import (
	"encoding/json"
	"strconv"
	"test/room"
	. "test/user"
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
		return
	}

	switch message.Type {
	case SubscribeEventType:
		roomNumber, err := strconv.ParseInt(message.Data, 10, 64)
		if err != nil {
			c.log.Error(err)
			return
		}
		conn.SetCurrentRoom(roomNumber)
		if c.rooms[roomNumber] == nil {
			newRoom := NewRoomConnections(roomNumber)
			c.rooms[roomNumber] = &newRoom
		}
		room := *c.rooms[roomNumber]
		room.AddConnection(conn)

		sendToFront := struct {
			RoomId int64
			User   *User
		}{
			RoomId: roomNumber,
			User:   conn.GetUser(),
		}
		bytes, err := json.Marshal(sendToFront)
		if err != nil {
			c.log.Error(err)
			return
		}
		message := EventMessage{
			Type: 1,
			Data: string(bytes),
		}
		room.SendMessage(message)
	case UnSubscribeRoomEventType:
		roomNumber, err := strconv.ParseInt(message.Data, 10, 64)
		if err != nil {
			c.log.Error(err)
			return
		}
		pRoom := c.rooms[roomNumber]
		if pRoom == nil {
			return
		}
		room := *pRoom
		room.RemoveConnection(conn)
		sendToFront := struct {
			RoomId int64
			UserId int64
		}{
			RoomId: roomNumber,
			UserId: conn.GetUser().Id,
		}
		bytes, err := json.Marshal(sendToFront)
		if err != nil {
			c.log.Error(err)
			return
		}
		message := EventMessage{
			Type: 2,
			Data: string(bytes),
		}
		room.SendMessage(message)
	case NewMessageEventType:
		message := room.Message{
			Text:      message.Data,
			Status:    1,
			RoomID:    conn.GetCurrentRoom(),
			CreatedBy: conn.GetUser().Id,
		}
		savedMessage, err := c.roomService.AddMessage(message)
		if err != nil {
			c.log.Error(err)
			return
		}
		c.SendMessageByRoom(conn.GetCurrentRoom(), savedMessage)
	}
}
