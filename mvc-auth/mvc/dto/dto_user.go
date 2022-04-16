package dto

type UserDto struct {
	Name string `Json:"name"`
	Id   int    `json:"id"`
}

type UsersDto []UserDto
