package service

import (
	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/MDmitryM/music-lib-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user musiclib.User) (uint, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (uint, error)
}

type Song interface {
	AddUserSong(userId uint, song musiclib.Song) (uint, error)
	GetUserSongs(userId uint, page, pageSize int) ([]musiclib.Song, error)
	GetUserSongById(userId uint, songId int) (musiclib.Song, error)
	UpdateUserSongInfo(userId uint, songId int, songInput musiclib.Song) (musiclib.Song, error)
}

type Service struct {
	Authorization
	Song
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Song:          NewSongService(repo),
	}
}
