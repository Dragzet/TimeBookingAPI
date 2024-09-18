package userModule

import (
	"TimeBookingAPI/internal/bookingModule"
	"context"
)

type UserStorage interface {
	Create(ctx context.Context, item *UserModel) error
	Find(ctx context.Context, username string) (*UserModel, error)
	Update(ctx context.Context, item *UserModel) error
	Delete(ctx context.Context, booking bookingModule.BookingStorage, username string) error
}
