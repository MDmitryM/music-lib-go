package repository

import (
	"errors"

	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/MDmitryM/music-lib-go/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SongPostgres struct {
	db *gorm.DB
}

func NewSongPostgres(db *gorm.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

func (r *SongPostgres) AddUserSong(userId uint, song musiclib.Song) (uint, error) {
	songModel := song.ToModel(userId)

	result := r.db.Create(&songModel)

	if result.Error != nil {
		logrus.Error(result.Error.Error())
		return 0, result.Error
	}

	return songModel.ID, nil
}

func (r *SongPostgres) GetUserSongs(userId uint, offset, pageSize int) ([]models.SongModel, error) {
	var songs []models.SongModel

	result := r.db.Where("user_id = ?", userId).
		Limit(pageSize).
		Offset(offset).
		Find(&songs)

	if result.Error != nil {
		return nil, result.Error
	}

	return songs, nil
}

func (r *SongPostgres) GetUserSongById(userId uint, songId int) (models.SongModel, error) {
	var songModel models.SongModel

	result := r.db.Where("user_id = ?", userId).
		Where("id = ?", songId).
		First(&songModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.SongModel{}, errors.New("song now found")
		}
		return models.SongModel{}, result.Error
	}

	return songModel, nil
}
