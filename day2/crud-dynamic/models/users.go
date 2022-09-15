package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `json:"id_user" form:"id_user"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserDTO struct {
	Name     string
	Email    string
	Password string
}

func ToUserDTO(user *User) *UserDTO {
	return &UserDTO{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func AssembUserDTO(user *UserDTO) *User {
	return &User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
