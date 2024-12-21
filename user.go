package musiclib

import "github.com/MDmitryM/music-lib-go/models"

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func (u *User) ToModel() models.UserModel {
	return models.UserModel{
		Email:        u.Email,
		PasswordHash: u.Password,
		Name:         u.Name,
	}
}
