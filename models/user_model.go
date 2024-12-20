package models

type UserModel struct {
	ID           uint        `gorm:"primaryKey"`
	Email        string      `gorm:"unique;not null"`
	PasswordHash string      `gorm:"not null"`
	Name         string      `gorm:"not null"`
	Songs        []SongModel `gorm:"foreignKey:UserID"`
}

func (UserModel) TableName() string {
	return "users"
}
