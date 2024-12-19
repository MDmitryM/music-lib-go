package service

import "github.com/MDmitryM/music-lib-go/pkg/repository"

type Authorization interface {
}

type Song interface {
}

type Service struct {
	Authorization
	Song
}

func NewService(repo *repository.Repository, signingkey, salt string) (*Service, error) {
	auth, err := NewAuthService(repo, signingkey, salt)
	if err != nil {
		return nil, err
	}
	return &Service{
		Authorization: auth,
		Song:          NewSongService(repo),
	}, nil
}
