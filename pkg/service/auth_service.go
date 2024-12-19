package service

import (
	"errors"

	"github.com/MDmitryM/music-lib-go/pkg/repository"
)

type AuthService struct {
	repo       *repository.Repository
	signingKey string
	salt       string
}

func NewAuthService(repo *repository.Repository, signingkey, salt string) (*AuthService, error) {
	if signingkey == "" || salt == "" {
		return nil, errors.New("invalid signing key or salt")
	}

	return &AuthService{repo: repo, signingKey: signingkey, salt: salt}, nil
}
