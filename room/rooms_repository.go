package room

import (
	"strings"
	"time"

	gm "gorm.io/gorm"
)

type RoomRepository interface {
	GetRoom(id int64) (*RoomsDto, []MessagesDto, error)
	CreateRoom(user RoomsDto) (*RoomsDto, error)
	UpdateRoom(user RoomsDto) (*RoomsDto, error)
	DeleteRoom(id int64) error
	GetRooms(filter *RoomsFilter) ([]RoomsDto, error)
}

type RoomRepositoryImpl struct {
	db *gm.DB
}

func (r *RoomRepositoryImpl) GetRoom(id int64) (*RoomsDto, []MessagesDto, error) {
	room := &RoomsDto{}
	messages := []MessagesDto{}
	if err := r.db.Where("id = ?", id).First(room).Error; err != nil {
		return nil, nil, err
	}
	if err := r.db.Where("Room_id = ?", id).Find(&messages).Error; err != nil {
		return nil, nil, err
	}
	return room, messages, nil
}

func (r *RoomRepositoryImpl) CreateRoom(room RoomsDto) (*RoomsDto, error) {
	room.CreatedAt = time.Now()
	if err := r.db.Create(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepositoryImpl) UpdateRoom(room RoomsDto) (*RoomsDto, error) {
	result, _, err := r.GetRoom(room.Id)
	if err != nil {
		return nil, err
	}
	room.CreatedAt = result.CreatedAt
	if err := r.db.Save(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepositoryImpl) DeleteRoom(id int64) error {
	if err := r.db.Delete(&RoomsDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoomRepositoryImpl) GetRooms(filter *RoomsFilter) (rooms []RoomsDto, err error) {
	var findResult *gm.DB = r.db
	var search []string
	if filter != nil {
		if filter.Name != nil {
			search = append(search, "name LIKE '%"+*filter.Name+"%'")
		}
		findResult = findResult.Where(strings.Join(search, " AND "))
	}
	if err := findResult.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return
}

func NewRoomRepository(db *gm.DB) RoomRepository {
	return &RoomRepositoryImpl{
		db,
	}
}
