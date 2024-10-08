package repository

import (
	"TimeBookingAPI/internal/storage/PostgreSQL"
)

type DB struct {
	Client PostgreSQL.Client
	UserStorage
	BookingStorage
}

func NewDB(client PostgreSQL.Client) *DB {
	return &DB{
		Client: client,
	}
}
