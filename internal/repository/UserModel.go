package repository

import (
	"time"
)

type UserModel struct {
	ID        string    `json:"ID" swaggerignore:"true"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
}

func NewUserModel() *UserModel {
	return &UserModel{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
