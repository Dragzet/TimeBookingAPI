package repository

import (
	"context"
)

type UserStorage interface {
	UserCreate(ctx context.Context, item *UserModel) error
	UserFind(ctx context.Context, username string) (*UserModel, error)
	UserUpdate(ctx context.Context, item *UserModel) error
	UserDelete(ctx context.Context, username string) error
}
