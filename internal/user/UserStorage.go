package user

import (
	"TimeBookingAPI/internal/booking"
	"context"
)

type UserStorage interface {
	Create(ctx context.Context, item *User) error
	Find(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, item *User) error
	Delete(ctx context.Context, booking booking.BookingStorage, id string) error
}
