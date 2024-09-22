package service

import (
	"TimeBookingAPI/internal/repository"
	"context"
	"fmt"
	"net/http"
	"time"
)

func (s Service) CreateBooking(ctx context.Context, booking *repository.BookingModel) (Answer, error) {
	answer := Answer{
		Status: http.StatusOK,
	}

	booking.EndTime = time.Now().Add(time.Hour * time.Duration(booking.Delta))

	if booking.Delta < 1 {
		answer.Status = http.StatusBadRequest
		return answer, fmt.Errorf("invalid booking data")
	}

	err := s.repo.BookingCreate(ctx, booking)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return answer, err
	}
	return answer, nil
}

func (s Service) DeleteBooking(ctx context.Context, id string) (Answer, error) {
	answer := Answer{
		Status: http.StatusOK,
	}

	if id == "" {
		answer.Status = http.StatusBadRequest
		return answer, fmt.Errorf("id is empty")
	}

	err := s.repo.BookingDelete(ctx, id)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return answer, err
	}
	return answer, nil
}

func (s Service) FindBooking(ctx context.Context, username string) (Answer, error) {
	answer := Answer{
		Status: http.StatusOK,
	}

	if username == "" {
		answer.Status = http.StatusBadRequest
		return answer, fmt.Errorf("username is empty")
	}

	bookings, err := s.repo.BookingFindAll(ctx, username)
	if err != nil {
		answer.Status = http.StatusInternalServerError
		return answer, err
	}

	answer.Data = bookings

	return answer, nil
}
