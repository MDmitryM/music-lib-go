package service

import (
	"encoding/json"

	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/MDmitryM/music-lib-go/pkg/repository"
)

type SongCacheService struct {
	repo repository.CacheSong
}

func NewSongCacheService(repo repository.CacheSong) *SongCacheService {
	return &SongCacheService{
		repo: repo,
	}
}

func (s *SongCacheService) CacheUserSong(userID, songID uint, song musiclib.Song) error {
	songJson, err := json.Marshal(song)
	if err != nil {
		return err
	}
	return s.repo.CacheUserSong(userID, songID, string(songJson))
}
