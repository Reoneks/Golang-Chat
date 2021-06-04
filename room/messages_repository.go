package room

import (
	"time"

	gm "gorm.io/gorm"
)

type MessagesRepository interface {
	CreateMessage(user MessagesDto) (*MessagesDto, error)
	UpdateMessage(user MessagesDto) (*MessagesDto, error)
	DeleteMessage(id int64) error
}

type MessagesRepositoryImpl struct {
	db *gm.DB
}

func (r *MessagesRepositoryImpl) CreateMessage(message MessagesDto) (*MessagesDto, error) {
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	if err := r.db.Create(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MessagesRepositoryImpl) UpdateMessage(message MessagesDto) (*MessagesDto, error) {
	var messages *MessagesDto
	if err := r.db.Where("id = ?", message.Id).First(&messages).Error; err != nil {
		return nil, err
	}
	message.CreatedAt = messages.CreatedAt
	message.UpdatedAt = time.Now()
	if err := r.db.Save(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MessagesRepositoryImpl) DeleteMessage(id int64) error {
	if err := r.db.Delete(&MessagesDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewMessagesRepository(db *gm.DB) MessagesRepository {
	return &MessagesRepositoryImpl{
		db,
	}
}
