package room

import "time"

type Messages struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	Status    int64     `json:"status"`
	RoomID    int64     `json:"room_id"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
