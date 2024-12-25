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

func (r *SongPostgres) GetUserSongById(userId uint, songId uint) (models.SongModel, error) {
	var songModel models.SongModel

	result := r.db.Where("user_id = ?", userId).
		Where("id = ?", songId).
		First(&songModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.SongModel{}, errors.New("song not found")
		}
		return models.SongModel{}, result.Error
	}

	return songModel, nil
}

func (r *SongPostgres) UpdateUserSongInfo(userId uint, songId uint, song musiclib.Song) (models.SongModel, error) {
	songModel := song.ToModel(userId)
	songModel.ID = songId

	result := r.db.Where("id = ? AND user_id = ?", songId, userId).Updates(&songModel)
	if result.Error != nil {
		return models.SongModel{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.SongModel{}, errors.New("song not found or user doesn't have access")
	}

	var updatedSong models.SongModel
	if err := r.db.First(&updatedSong, songId).Error; err != nil {
		return models.SongModel{}, err
	}

	return updatedSong, nil
}

func (r *SongPostgres) DeleteUserSongByID(userId, songId uint) error {
	result := r.db.Where("id = ? AND user_id = ?", songId, userId).Delete(&models.SongModel{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("song not found or user doesn't have access")
	}

	return nil
}
