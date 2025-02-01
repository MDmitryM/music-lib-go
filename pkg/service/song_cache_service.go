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

func (s *SongCacheService) GetUserCachedSongByID(userID, songID uint) (musiclib.Song, error) {
	data, err := s.repo.GetUserCachedSongByID(userID, songID)
	if err != nil {
		return musiclib.Song{}, err
	}

	var cachedSong musiclib.Song
	if err := json.Unmarshal([]byte(data), &cachedSong); err != nil {
		return musiclib.Song{}, err
	}

	return cachedSong, nil
}

func (s *SongCacheService) DeleteUserCachedSong(userID, songID uint) error {
	return s.repo.DeleteUserCachedSong(userID, songID)
}
