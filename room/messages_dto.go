package room

import "time"

type MessagesDto struct {
	Id        int64     `gorm:"column:id"`
	Text      string    `gorm:"column:text"`
	Status    int64     `gorm:"column:status"`
	RoomID    int64     `gorm:"column:room_id"`
	CreatedBy int64     `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (MessagesDto) TableName() string {
	return "messages"
}
