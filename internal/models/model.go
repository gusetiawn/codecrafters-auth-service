package models

import "time"

type User struct {
	ID          int64      `db:"id"`
	Username    string     `db:"username"`
	Password    string     `db:"password"`
	Email       string     `db:"email"`
	FirstName   string     `db:"first_name"`
	LastName    string     `db:"last_name"`
	DateOfBirth time.Time  `db:"date_of_birth"`
	Gender      string     `db:"gender"`
	PhoneNumber string     `db:"phone_number"`
	Address     string     `db:"address"`
	IsActive    bool       `db:"is_active"`
	CreatedAt   *time.Time `db:"created_at"`
}

type UserRole struct {
	ID     int `db:"user_role_id"`
	UserID int `db:"user_id"`
	RoleID int `db:"role_id"`
}

type Role struct {
	ID          int    `db:"role_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type Token struct {
	ID         int       `db:"token_id"`
	UserID     int       `db:"user_id"`
	Token      string    `db:"token"`
	Expiration time.Time `db:"expiration"`
	CreatedAt  time.Time `db:"created_at"`
}

type ActivityLog struct {
	ID           int       `db:"log_id"`
	UserID       int       `db:"user_id"`
	ActivityType string    `db:"activity_type"`
	Timestamp    time.Time `db:"timestamp"`
	IPAddress    string    `db:"ip_address"`
}
