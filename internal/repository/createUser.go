package repository

import (
	"context"
	"fmt"
	"github.com/gusetiawn/codecrafters-auth-service/internal/models"
)

const QUERY_REGISTER = `INSERT INTO users (username, password, email, first_name, last_name, date_of_birth, gender, phone_number, address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

func (repo *repository) CreateUser(ctx context.Context, req models.Register) error {

	_, err := repo.db.ExecContext(ctx, QUERY_REGISTER, req.Username, req.Password, req.Email, req.FirstName, req.LastName,
		req.DateOfBirth, req.Gender, req.PhoneNumber, req.Address)
	if err != nil {
		fmt.Println("[repository][func: CreateUser][error]: ", err)
		return err
	}

	return nil
}
