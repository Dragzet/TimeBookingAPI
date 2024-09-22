package repository

import (
	"context"
)

type BookingStorage interface {
	BookingCreate(ctx context.Context, item *BookingModel) error
	BookingFind(ctx context.Context, username string) (*BookingModel, error)
	BookingFindAll(ctx context.Context, username string) ([]*BookingModel, error)
	BookingDelete(ctx context.Context, id string) error
}
