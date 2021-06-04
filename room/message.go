package room

import "time"

type StatusType int64

//^ Status types
const (
	Unread StatusType = 1
	Read   StatusType = 2
)

func (s StatusType) ToInt64() int64 {
	return int64(s)
}

func NewStatusType(i int64) StatusType {
	return StatusType(i)
}

type Message struct {
	Id        int64      `json:"id"`
	Text      string     `json:"text"`
	Status    StatusType `json:"status"`
	RoomID    int64      `json:"room_id"`
	CreatedBy int64      `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
