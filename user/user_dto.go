package user

type UserDto struct {
	Id        int64  `gorm:"column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	Status    int64  `gorm:"column:status"`
}

func (UserDto) TableName() string {
	return "users"
}
