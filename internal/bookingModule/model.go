package bookingModule

import "time"

type BookingModel struct {
	ID        string
	UserID    string
	StartTime time.Time
	EndTime   time.Time
}

func New(userID string, delta int) *BookingModel {
	return &BookingModel{
		UserID:    userID,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Duration(delta) * time.Hour),
	}
}
