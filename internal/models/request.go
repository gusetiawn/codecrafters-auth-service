package models

import "time"

type Register struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUser struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserInfo struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
}
