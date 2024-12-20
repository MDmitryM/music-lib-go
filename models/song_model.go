package models

type SongModel struct {
	ID     uint   `gorm:"primaryKey"`
	Artist string `gorm:"not null"`
	Title  string `gorm:"not null"`
	Album  string
	Year   int
	UserID uint `gorm:"not null"`
}

func (SongModel) TableName() string {
	return "songs"
}
