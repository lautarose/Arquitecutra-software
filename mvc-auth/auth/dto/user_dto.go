package dto

type UserDto struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

type UsersDto []UserDto
