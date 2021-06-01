package user

import (
	"errors"
	"test/utils"
)

type UserService interface {
	GetUser(userId int64) (*User, error)
	GetUserByEmail(username string) (*User, error)
	Registration(user User) (*User, error)
	GetUsers(filter *UserFilter) ([]User, error)
	DeleteUser(user int64) error
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func (s *UserServiceImpl) GetUser(userId int64) (*User, error) {
	result, err := s.userRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	resultUser := FromUserDto(*result)
	return &resultUser, nil
}

func (s *UserServiceImpl) Registration(user User) (*User, error) {
	password, err := utils.Encrypt(user.Password)
	user.Password = password
	if err != nil {
		return nil, err
	}
	user.Status = Active
	result, err := s.userRepository.CreateUser(ToUserDto(user))
	if err != nil {
		return nil, err
	}
	resultUser := FromUserDto(*result)
	return &resultUser, nil
}
func (s *UserServiceImpl) GetUsers(filter *UserFilter) ([]User, error) {
	result, err := s.userRepository.GetUsers(filter)
	if err != nil {
		return nil, err
	}
	resultUser := FromUserDtos(result)
	return resultUser, nil
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*User, error) {
	userStatus := []int64{Active.ToInt64()}
	filter := &UserFilter{
		Email:  &email,
		Status: userStatus,
	}
	someUsers, err := s.userRepository.GetUsers(filter)
	if err != nil {
		return nil, err
	} else if len(someUsers) == 0 {
		return nil, errors.New("user not found")
	}
	resultUser := FromUserDto(someUsers[0])
	return &resultUser, nil
}

func (s *UserServiceImpl) DeleteUser(user int64) error {
	return s.userRepository.DeleteUser(user)
}

func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{
		userRepository,
	}
}
