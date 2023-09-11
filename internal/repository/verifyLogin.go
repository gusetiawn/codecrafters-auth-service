package repository

import (
	"context"
	"database/sql"
	"fmt"
)

const QUERY_LOGIN = `SELECT password FROM users WHERE username = $1`

func (repo *repository) VerifyLogin(ctx context.Context, username string, password string) (bool, error) {
	var storedPassword string

	fmt.Println("username: ", username)
	fmt.Println("query: ", QUERY_LOGIN)
	err := repo.db.
		QueryRowContext(ctx, QUERY_LOGIN, username).
		Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, fmt.Errorf("failed executing verify login: %s", err.Error())
	}
	if storedPassword != password {
		return false, fmt.Errorf("unauthorized")
	}

	return true, nil
}
