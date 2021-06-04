package room

import (
	"errors"
)

type RoomService interface {
	GetRoom(id int64) (*Rooms, []Message, error)
	GetRooms(filter *RoomsFilter) ([]Rooms, error)
	CreateRoom(room Rooms) (*Rooms, error)
	DeleteRoom(room_id, userId int64) error
	UpdateRoom(room Rooms, userId int64) (*Rooms, error)
	AddUsers(roomId, userId int64, users []int64) (errorsArray error)
	AddMessage(message Message) (*Message, error)
	UpdateMessage(message Message) (*Message, error)
	DeleteMessage(messageId int64) error
}

type RoomServiceImpl struct {
	RoomRepository     RoomRepository
	uRRepository       RoomUsersRepository
	messagesRepository MessagesRepository
}

func (s *RoomServiceImpl) GetRoom(id int64) (*Rooms, []Message, error) {
	result, Messages, err := s.RoomRepository.GetRoom(id)
	if err != nil {
		return nil, nil, err
	}
	resultRoom := FromRoomsDto(*result)
	resultMessages := FromMessagesDtos(Messages)
	return &resultRoom, resultMessages, nil
}

func (s *RoomServiceImpl) GetRooms(filter *RoomsFilter) ([]Rooms, error) {
	result, err := s.RoomRepository.GetRooms(filter)
	if err != nil {
		return nil, err
	}
	resultRooms := FromRoomsDtos(result)
	return resultRooms, nil
}

func (s *RoomServiceImpl) CreateRoom(room Rooms) (*Rooms, error) {
	result, err := s.RoomRepository.CreateRoom(ToRoomsDto(room))
	if err != nil {
		return nil, err
	}
	_, err = s.uRRepository.CreateUserRoomConnection(RoomUsersDto{
		UserID: result.CreatedBy,
		RoomID: result.Id,
	})
	if err != nil {
		return nil, err
	}
	resultMessages := FromRoomsDto(*result)
	return &resultMessages, nil
}

func (s *RoomServiceImpl) DeleteRoom(Room_id, userId int64) error {
	result, _, err := s.GetRoom(Room_id)
	if err != nil {
		return err
	} else if result.CreatedBy != userId {
		return errors.New("you are not allowed to do it")
	}
	return s.RoomRepository.DeleteRoom(Room_id)
}

func (s *RoomServiceImpl) UpdateRoom(room Rooms, userId int64) (*Rooms, error) {
	result, _, err := s.GetRoom(room.Id)
	if err != nil {
		return nil, err
	} else if result.CreatedBy != userId {
		return nil, errors.New("you are not allowed to do it")
	}
	updateResult, err := s.RoomRepository.UpdateRoom(ToRoomsDto(room))
	if err != nil {
		return nil, err
	}
	resultRoom := FromRoomsDto(*updateResult)
	return &resultRoom, nil
}

func (s *RoomServiceImpl) AddUsers(roomId, userId int64, users []int64) (err error) {
	result, _, err := s.GetRoom(roomId)
	if err != nil {
		return
	} else if result.CreatedBy != userId {
		err = errors.New("you are not allowed to do it")
		return
	}
	for _, user := range users {
		_, err = s.uRRepository.CreateUserRoomConnection(RoomUsersDto{
			UserID: user,
			RoomID: roomId,
		})
		if err != nil {
			return
		}
	}
	return
}

func (s *RoomServiceImpl) AddMessage(message Message) (*Message, error) {
	result, err := s.messagesRepository.CreateMessage(ToMessagesDto(message))
	if err != nil {
		return nil, err
	}
	resultMessages := FromMessagesDto(*result)
	return &resultMessages, nil
}

func (s *RoomServiceImpl) UpdateMessage(message Message) (*Message, error) {
	result, err := s.messagesRepository.UpdateMessage(ToMessagesDto(message))
	if err != nil {
		return nil, err
	}
	resultMessages := FromMessagesDto(*result)
	return &resultMessages, nil
}

func (s *RoomServiceImpl) DeleteMessage(messageId int64) error {
	return s.messagesRepository.DeleteMessage(messageId)
}

func NewRoomService(RoomRepository RoomRepository, uRRepository RoomUsersRepository, messagesRepository MessagesRepository) RoomService {
	return &RoomServiceImpl{
		RoomRepository,
		uRRepository,
		messagesRepository,
	}
}
