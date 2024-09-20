package bookingModule

import "time"

// BookingModel представляет собой модель бронирования в системе.
type BookingModel struct {
	ID        string    `json:"id"`        // Уникальный идентификатор бронирования
	Username  string    `json:"username"`  // Имя пользователя, сделавшего бронирование
	StartTime time.Time `json:"startTime"` // Время начала бронирования
	EndTime   time.Time `json:"endTime"`   // Время окончания бронирования
	Delta     int       `json:"delta"`     // Длительность бронирования (например, в минутах)
}

// New создает и возвращает новую модель бронирования с текущим временем начала.
func New() *BookingModel {
	return &BookingModel{
		StartTime: time.Now(), // Устанавливаем время начала бронирования
	}
}
