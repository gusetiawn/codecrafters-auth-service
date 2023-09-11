package repository

//go:generate mockgen -source=./repository.go -destination=./mock/mock.go -package=mock

import (
	"context"
	"database/sql"
	"github.com/gusetiawn/codecrafters-auth-service/internal/models"
)

type Repository interface {
	CreateUser(ctx context.Context, req models.Register) error
	VerifyLogin(ctx context.Context, username string, password string) (bool, error)
	GetUser(ctx context.Context, username string) (*models.UserInfo, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
