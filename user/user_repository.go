package user

import (
	"fmt"
	"strings"

	gm "gorm.io/gorm"
)

type UserRepository interface {
	GetUser(id int64) (*UserDto, error)
	CreateUser(user UserDto) (*UserDto, error)
	UpdateUser(user UserDto) (*UserDto, error)
	DeleteUser(id int64) error
	GetUsers(filter *UserFilter) ([]UserDto, error)
}

type UserRepositoryImpl struct {
	db *gm.DB
}

func (r *UserRepositoryImpl) GetUser(id int64) (*UserDto, error) {
	user := &UserDto{}
	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) CreateUser(user UserDto) (*UserDto, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user UserDto) (*UserDto, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id int64) error {
	if err := r.db.Delete(&UserDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) GetUsers(filter *UserFilter) (users []UserDto, err error) {
	var findResult *gm.DB = r.db
	var search []string
	if filter != nil {
		if filter.Email != nil {
			search = append(search, "email LIKE '%"+*filter.Email+"%'")
		}
		if filter.FirstName != nil {
			search = append(search, "first_name LIKE '%"+*filter.FirstName+"%'")
		}
		if filter.LastName != nil {
			search = append(search, "last_name LIKE '%"+*filter.LastName+"%'")
		}
		if filter.Status != nil && len(filter.Status) > 0 {
			search = append(
				search,
				fmt.Sprintf(
					"status IN (%s)",
					strings.Trim(strings.Replace(fmt.Sprint(filter.Status), " ", ",", -1), "[]"),
				),
			)
		}
	}
	findResult = findResult.Where(strings.Join(search, " AND "))
	if err = findResult.Find(&users).Error; err != nil {
		return nil, err
	}
	return
}

func NewUserRepository(db *gm.DB) UserRepository {
	return &UserRepositoryImpl{
		db,
	}
}
