package musiclib

import "github.com/MDmitryM/music-lib-go/models"

type Song struct {
	Artist string `json:"artist" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Album  string `json:"album"`
	Year   string `json:"year"`
}

func (s *Song) ToModel(userId uint) models.SongModel {
	return models.SongModel{
		Artist: s.Artist,
		Title:  s.Title,
		Album:  s.Album,
		Year:   s.Year,
		UserID: userId,
	}
}

func FromModel(song models.SongModel) Song {
	return Song{
		Artist: song.Artist,
		Title:  song.Title,
		Album:  song.Album,
		Year:   song.Year,
	}
}
