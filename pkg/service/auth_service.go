package service

import (
	"os"
	"time"

	musiclib "github.com/MDmitryM/music-lib-go"
	"github.com/MDmitryM/music-lib-go/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.Repository
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"id"`
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(input musiclib.User) (uint, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	input.Password = string(hash)

	return s.repo.Authorization.CreateUser(input)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	id, err := s.repo.Authorization.IsUserValid(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
