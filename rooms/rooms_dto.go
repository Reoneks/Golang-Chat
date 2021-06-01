package rooms

import (
	"time"
)

type RoomsDto struct {
	Id        int64     `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Status    int64     `gorm:"column:status"`
	CreatedBy int64     `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (RoomsDto) TableName() string {
	return "rooms"
}
