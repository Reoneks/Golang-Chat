package rooms

import (
	gm "gorm.io/gorm"
)

type RoomUsersRepository interface {
	GetRoomsByUserId(id int64) ([]RoomsDto, error)
	CreateUserRoomConnection(user RoomUsersDto) (*RoomUsersDto, error)
}

type RoomUsersRepositoryImpl struct {
	db *gm.DB
}

func (r *RoomUsersRepositoryImpl) GetRoomsByUserId(id int64) (rooms []RoomsDto, err error) {
	var roomUsers []RoomUsersDto
	if err = r.db.Where("user_id = ?", id).Find(&roomUsers).Error; err != nil {
		return
	}
	for _, user := range roomUsers {
		var roomStruct *RoomsDto
		if err = r.db.Where("id = ?", user.RoomID).First(&roomStruct).Error; err != nil {
			rooms = nil
			return
		}
		rooms = append(rooms, *roomStruct)
	}
	return
}

func (r *RoomUsersRepositoryImpl) CreateUserRoomConnection(user RoomUsersDto) (*RoomUsersDto, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewRoomUsersRepository(db *gm.DB) RoomUsersRepository {
	return &RoomUsersRepositoryImpl{
		db,
	}
}
