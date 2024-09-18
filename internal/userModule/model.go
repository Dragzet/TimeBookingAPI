package userModule

import (
	"time"
)

type UserModel struct {
	ID        string
	Username  string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(username string, password []byte) *UserModel {
	return &UserModel{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
