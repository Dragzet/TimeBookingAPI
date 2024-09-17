package booking

import (
	"context"
)

type BookingStorage interface {
	Create(ctx context.Context, item *Booking) error
	Find(ctx context.Context, id string) (*Booking, error)
	FindAll(ctx context.Context, userID string) ([]*Booking, error)
	Delete(ctx context.Context, id string) error
}
