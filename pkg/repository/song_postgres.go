package repository

import "gorm.io/gorm"

type SongPostgres struct {
	db *gorm.DB
}

func NewSongPostgres(db *gorm.DB) *SongPostgres {
	return &SongPostgres{db: db}
}
