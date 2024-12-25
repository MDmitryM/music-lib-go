package service

import (
	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/MDmitryM/music-lib-go/pkg/repository"
)

type SongService struct {
	repo *repository.Repository
}

func NewSongService(repo *repository.Repository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) AddUserSong(userId uint, song musiclib.Song) (uint, error) {
	return s.repo.Song.AddUserSong(userId, song)
}

func (s *SongService) GetUserSongs(userId uint, page, pageSize int) ([]musiclib.Song, error) {
	offset := (page - 1) * pageSize
	songsModel, err := s.repo.Song.GetUserSongs(userId, offset, pageSize)
	if err != nil {
		return nil, err
	}

	var songs []musiclib.Song
	for _, v := range songsModel {
		songs = append(songs, musiclib.FromModel(v))
	}

	return songs, nil
}

func (s *SongService) GetUserSongById(userId uint, songId int) (musiclib.Song, error) {
	songModel, err := s.repo.Song.GetUserSongById(userId, uint(songId))
	if err != nil {
		return musiclib.Song{}, err
	}
	song := musiclib.FromModel(songModel)

	return song, nil
}

func (s *SongService) UpdateUserSongInfo(userId uint, songId int, songInput musiclib.Song) (musiclib.Song, error) {
	updatedSongModel, err := s.repo.Song.UpdateUserSongInfo(userId, uint(songId), songInput)
	if err != nil {
		return musiclib.Song{}, err
	}

	updatedSong := musiclib.FromModel(updatedSongModel)

	return updatedSong, nil
}

func (s *SongService) DeleteUserSongByID(userId uint, songId int) error {
	return s.repo.Song.DeleteUserSongByID(userId, uint(songId))
}
