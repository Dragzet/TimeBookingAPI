package user

import (
	"time"
)

type User struct {
	ID        string
	Username  string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(username string, password []byte) *User {
	return &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
