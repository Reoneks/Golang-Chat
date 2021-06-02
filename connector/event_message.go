package connector

type EventType int64

const (
	SubscribeEventType       EventType = 1 // Connect to room		Data: "1","2","3"
	UnSubscribeRoomEventType EventType = 2 // Disconnect from room  Data: ""
	NewMessageEventType      EventType = 3 // Message				Data: "Text Message"
)

type EventMessage struct {
	Type EventType `json:"type"`
	Data string    `json:"Data"`
}
