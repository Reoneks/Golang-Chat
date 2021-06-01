package user

func FromUserDto(userDto UserDto) User {
	return User{
		Id:        userDto.Id,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
		Password:  userDto.Password,
		Status:    NewStatusCode(userDto.Status),
	}
}

func FromUserDtos(UserDtos []UserDto) (users []User) {
	for _, dto := range UserDtos {
		users = append(users, FromUserDto(dto))
	}
	return
}

func ToUserDto(user User) UserDto {
	return UserDto{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status.ToInt64(),
	}
}

func ToUserDtos(users []User) (userDtos []UserDto) {
	for _, dto := range users {
		userDtos = append(userDtos, ToUserDto(dto))
	}
	return
}
