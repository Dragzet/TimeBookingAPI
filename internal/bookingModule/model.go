package bookingModule

import "time"

type BookingModel struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func New(userID string, delta int) *BookingModel {
	return &BookingModel{
		UserID:    userID,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Duration(delta) * time.Hour),
	}
}
