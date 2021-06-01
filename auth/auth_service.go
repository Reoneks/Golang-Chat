package auth

import (
	"test/user"
	"test/utils"

	"github.com/go-chi/jwtauth"
)

type Login struct {
	Token string     `json:"token"`
	User  *user.User `json:"user"`
}

type AuthService interface {
	Login(email, password string) (*Login, error)
}

type AuthServiceImpl struct {
	jwt         *jwtauth.JWTAuth
	userService user.UserService
}

func (s *AuthServiceImpl) Login(email, password string) (*Login, error) {
	receivedUser, err := s.userService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	err = utils.Compare(receivedUser.Password, password)
	if err != nil {
		return nil, err
	}
	_, tokenString, err := s.jwt.Encode(map[string]interface{}{"user_id": receivedUser.Id})
	if err != nil {
		return nil, err
	}
	return &Login{
		Token: tokenString,
		User:  receivedUser,
	}, nil
}

func NewAuthService(userService user.UserService, jwt *jwtauth.JWTAuth) AuthService {
	return &AuthServiceImpl{
		jwt,
		userService,
	}
}
