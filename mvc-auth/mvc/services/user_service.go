package services

import (
	userCliente "mvc-auth/mvc/clients/user"
	"mvc-auth/mvc/dto"
	"mvc-auth/mvc/model"
	e "mvc-auth/mvc/utils"
)

type userService struct{}

type userServiceInterface interface {
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.Name = userDto.Name

	user = userCliente.InsertUser(user)

	return userDto, nil
}
