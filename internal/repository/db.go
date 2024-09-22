package repository

import (
	"TimeBookingAPI/internal/storage/PostgreSQL"
)
//go:generate go run github.com/vektra/mockery/v2@v2.46 --name UserStorage
//go:generate go run github.com/vektra/mockery/v2@v2.46 --name BookingStorage
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
