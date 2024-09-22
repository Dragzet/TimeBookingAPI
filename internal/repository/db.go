package repository

import (
	"TimeBookingAPI/internal/storage/PostgreSQL"
)

type DB struct {
	client PostgreSQL.Client
	UserStorage
	BookingStorage
}

func NewDB(client PostgreSQL.Client) *DB {
	return &DB{
		client: client,
	}
}
