package service

import "github.com/MDmitryM/music-lib-go/pkg/repository"

type SongService struct {
	repo *repository.Repository
}

func NewSongService(repo *repository.Repository) *SongService {
	return &SongService{repo: repo}
}
