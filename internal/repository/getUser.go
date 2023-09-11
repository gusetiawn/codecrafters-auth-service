package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gusetiawn/codecrafters-auth-service/internal/models"
)

const QUERY_GET_USER = `SELECT id, username, email, first_name, last_name, date_of_birth, gender, phone_number, address FROM users WHERE username = $1`

func (repo *repository) GetUser(ctx context.Context, username string) (*models.UserInfo, error) {
	var result models.UserInfo

	err := repo.db.
		QueryRowContext(ctx, QUERY_GET_USER, username).
		Scan(&result.ID, &result.Username, &result.Email, &result.FirstName, &result.LastName,
			&result.DateOfBirth, &result.Gender, &result.PhoneNumber, &result.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return &result, nil
		}

		return &result, fmt.Errorf("failed executing get user: %s", err.Error())
	}

	return &result, nil
}
