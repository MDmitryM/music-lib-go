package repository

import (
	musiclib "github.com/MDmitryM/music-lib-go"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(input musiclib.User) (uint, error)
	IsUserValid(email, password string) (uint, error)
}

type Song interface {
}

type Repository struct {
	Authorization
	Song
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Song:          NewSongPostgres(db),
	}
}
