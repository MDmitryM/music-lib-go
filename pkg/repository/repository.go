package repository

import (
	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/MDmitryM/music-lib-go/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(input musiclib.User) (uint, error)
	IsUserValid(email, password string) (uint, error)
}

type Song interface {
	AddUserSong(userId uint, song musiclib.Song) (uint, error)
	GetUserSongs(userId uint, offset, pageSize int) ([]models.SongModel, error)
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
