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
