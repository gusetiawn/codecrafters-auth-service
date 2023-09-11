package service

import (
	"context"
	"fmt"
	"github.com/gusetiawn/codecrafters-auth-service/internal/jwt"
	"github.com/gusetiawn/codecrafters-auth-service/internal/models"
)

func (s *service) Login(ctx context.Context, req models.Login) (string, error) {

	result, err := s.repository.VerifyLogin(ctx, req.Username, req.Password)
	if err != nil {
		fmt.Println("[service][func: Login][error]: ", err)
		return "", err
	}

	if !result {
		return "", fmt.Errorf("unathorized")
	}
	token, err := jwt.GenerateToken(req.Username)
	if err != nil {
		return "", err
	}
	fmt.Println(token)

	return token, nil
}
