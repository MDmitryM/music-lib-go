package repository

import "gorm.io/gorm"

type Authorization interface {
}

type Song interface {
}

type Repository struct {
	Authorization
	Song
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Song:          NewSongPostgres(db),
	}
}
