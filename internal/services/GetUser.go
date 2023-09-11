package service

import (
	"context"
	"fmt"
	"github.com/gusetiawn/codecrafters-auth-service/internal/jwt"
	"github.com/gusetiawn/codecrafters-auth-service/internal/models"
)

func (s *service) GetUser(ctx context.Context, token string, username string) (*models.UserInfo, error) {

	result, err := s.repository.GetUser(ctx, username)
	if err != nil {
		fmt.Println("[service][func: Login][error]: ", err)
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf("unathorized")
	}
	claim, err := jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	fmt.Println(claim)
	if claim.Username != username {
		return nil, fmt.Errorf("unAuthorized")
	}

	return result, nil
}
