package service

import (
	"TimeBookingAPI/internal/repository"
	"context"
	"encoding/json"
)

//go:generate mockgen -source=service.go -destination=mock/mock.go
type Answer struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (answer *Answer) GetJson() []byte {
	d, err := json.Marshal(answer)
	if err != nil {
		return []byte{}
	}
	return d
}

type UserService interface {
	CreateUser(ctx context.Context, user *repository.UserModel) (Answer, error)
	DeleteUser(ctx context.Context, username string) (Answer, error)
}

type BookingService interface {
	CreateBooking(ctx context.Context, booking *repository.BookingModel) (Answer, error)
	DeleteBooking(ctx context.Context, id string) (Answer, error)
	FindBooking(ctx context.Context, username string) (Answer, error)
}

type Service struct {
	repo           *repository.DB
	UserService    UserService
	BookingService BookingService
}

func NewService(repo *repository.DB) *Service {
	return &Service{
		repo: repo,
	}
}
