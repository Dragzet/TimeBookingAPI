package repository

import "time"

type BookingModel struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Delta     int       `json:"delta"`
}

func NewBookingModel() *BookingModel {
	return &BookingModel{
		StartTime: time.Now(),
	}
}
