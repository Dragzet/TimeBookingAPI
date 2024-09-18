package userModule

import (
	"time"
)

type UserModel struct {
	ID        string    `json:"ID"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func New() *UserModel {
	return &UserModel{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
