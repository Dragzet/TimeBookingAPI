package user

import (
	"time"
)

type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(username, password string) *User {
	return &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
