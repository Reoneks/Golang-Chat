package rooms

type RoomUsersDto struct {
	UserID int64 `gorm:"column:user_id"`
	RoomID int64 `gorm:"column:room_id"`
}

func (RoomUsersDto) TableName() string {
	return "room_users"
}
