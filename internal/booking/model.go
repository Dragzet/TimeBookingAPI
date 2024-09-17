package booking

import "time"

type Booking struct {
	ID        string
	UserID    string
	StartTime time.Time
	EndTime   time.Time
}

func New(userID string, delta int) *Booking {
	return &Booking{
		UserID:    userID,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Duration(delta) * time.Hour),
	}
}
