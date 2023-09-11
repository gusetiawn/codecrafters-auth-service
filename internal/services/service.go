package service

//go:generate mockgen -source=./service.go -destination=./mock/mock.go -package=mock

import (
	"context"
	"github.com/gusetiawn/codecrafters-auth-service/internal/models"
	"github.com/gusetiawn/codecrafters-auth-service/internal/repository"
)

type Service interface {
	Register(ctx context.Context, req models.Register) error
	Login(ctx context.Context, req models.Login) (string, error)
	GetUser(ctx context.Context, token string, username string) (*models.UserInfo, error)
}

type service struct {
	repository repository.Repository
}

func New(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}
