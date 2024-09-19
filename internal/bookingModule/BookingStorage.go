package bookingModule

import (
	"context"
)

type BookingStorage interface {
	Create(ctx context.Context, item *BookingModel) error
	Find(ctx context.Context, username string) (*BookingModel, error)
	FindAll(ctx context.Context, username string) ([]*BookingModel, error)
	Delete(ctx context.Context, id string) error
}
