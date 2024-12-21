package repository

import (
	"errors"

	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/MDmitryM/music-lib-go/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(input musiclib.User) (uint, error) {
	userModel := input.ToModel()

	result := r.db.Create(&userModel)

	if result.Error != nil {
		return 0, result.Error
	}

	return userModel.ID, nil
}

func (r *AuthPostgres) IsUserValid(email, password string) (uint, error) {
	var user models.UserModel

	result := r.db.Where("email=?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("email not found")
		}
		return 0, result.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return 0, errors.New("password is incorrect")
	}

	return user.ID, nil
}
