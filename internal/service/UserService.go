package service

import (
	"TimeBookingAPI/internal/repository"
	"context"
	"fmt"
	"net/http"
)

func (s Service) CreateUser(ctx context.Context, user *repository.UserModel) (Answer, error) {
	answer := Answer{
		Status: http.StatusOK,
	}

	if len(user.Username) < 5 || len(user.Password) < 5 || len(user.Password) > 72 {
		answer.Status = http.StatusBadRequest
		return answer, fmt.Errorf("invalid user data")
	}

	err := s.repo.UserCreate(ctx, user)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return answer, err
	}
	return answer, nil
}

func (s Service) DeleteUser(ctx context.Context, username string) (Answer, error) {
	answer := Answer{
		Status: http.StatusOK,
	}

	if username == "" {
		answer.Status = http.StatusBadRequest
		return answer, fmt.Errorf("invalid username")
	}

	err := s.repo.UserDelete(ctx, username)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return answer, err
	}
	return answer, nil
}

func (s Service) TestUserCreate(ctx context.Context, user *repository.UserModel) (Answer, error) {
	answer := Answer{
		Status: http.StatusOK,
	}

	if len(user.Username) < 5 || len(user.Password) < 5 || len(user.Password) > 72 {
		answer.Status = http.StatusBadRequest
		return answer, fmt.Errorf("invalid user data")
	}

	err := s.repo.TestUserCreate(ctx, user)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return answer, err
	}
	return answer, nil
}
