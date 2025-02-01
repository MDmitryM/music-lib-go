package repository

import (
	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/MDmitryM/music-lib-go/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(input musiclib.User) (uint, error)
	IsUserValid(email, password string) (uint, error)
}

type Song interface {
	AddUserSong(userId uint, song musiclib.Song) (uint, error)
	GetUserSongs(userId uint, offset, pageSize int) ([]models.SongModel, error)
	GetUserSongById(userId, songId uint) (models.SongModel, error)
	UpdateUserSongInfo(userId, songId uint, song musiclib.Song) (models.SongModel, error)
	DeleteUserSongByID(userId, songId uint) error
}

type CacheSong interface {
	CacheUserSong(userID, songID uint, data string) error
}

type Repository struct {
	Authorization
	Song
	CacheSong
}

func NewRepository(db *gorm.DB, redis *RedisDB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Song:          NewSongPostgres(db),
		CacheSong:     NewSongRedis(redis),
	}
}
