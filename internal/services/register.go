package service

import (
	"context"
	"fmt"
	"github.com/gusetiawn/codecrafters-auth-service/internal/models"
)

func (s *service) Register(ctx context.Context, req models.Register) error {

	if len(req.Password) < 12 {
		return fmt.Errorf("password must more than 12 characters")
	}
	err := s.repository.CreateUser(ctx, req)
	if err != nil {
		fmt.Println("[service][func: Register][error]: ", err)
	}

	return nil
}
