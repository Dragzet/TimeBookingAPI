package repository

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const userDBErrorStatement = "internal/userModule/UserDB.go: "

func (d DB) UserCreate(ctx context.Context, user *UserModel) error {
	stmt := `
	INSERT INTO users
		(username, password, created_at, updated_at)
	VALUES 
	    ($1, $2, $3, $4)
	RETURNING id
	`

	_, err := d.UserFind(ctx, user.Username)
	if err == nil {
		return fmt.Errorf("%s create userModule: user already exists", userDBErrorStatement)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s create hashing userModule: %s", userDBErrorStatement, err.Error())
	}

	if err := d.Client.QueryRow(ctx, stmt, user.Username, hash, user.CreatedAt, user.UpdatedAt).Scan(&user.ID); err != nil {
		return fmt.Errorf("%s create userModule: %s", userDBErrorStatement, err.Error())
	}
	return nil
}

func (d DB) UserFind(ctx context.Context, username string) (*UserModel, error) {
	stmt := `
		SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1
	`
	var user UserModel
	err := d.Client.QueryRow(ctx, stmt, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s find userModule: %s", userDBErrorStatement, err.Error())
	}
	return &user, nil
}

func (d DB) UserDelete(ctx context.Context, username string) error {
	stmt := `
		DELETE FROM users WHERE username = $1
	`

	user, err := d.UserFind(ctx, username)
	if err != nil {
		return fmt.Errorf("%s delete userModule findUser: %s", userDBErrorStatement, err.Error())
	}

	bookingsArr, err := d.BookingFindAll(ctx, user.Username)
	if err != nil {
		return fmt.Errorf("%s delete userModule findAll: %s", userDBErrorStatement, err.Error())
	}
	for _, tempBooking := range bookingsArr {
		if err := d.BookingDelete(ctx, tempBooking.ID); err != nil {
			return fmt.Errorf("%s delete userModule DeleteBooking: %s", userDBErrorStatement, err.Error())
		}
	}

	if _, err := d.Client.Exec(ctx, stmt, user.Username); err != nil {
		return fmt.Errorf("%s delete userModule: %s", userDBErrorStatement, err.Error())
	}
	return nil
}

// create user without hashing password for tests
func (d DB) TestUserCreate(ctx context.Context, user *UserModel) error {
	stmt := `
	INSERT INTO users
		(username, password, created_at, updated_at)
	VALUES 
	    ($1, $2, $3, $4)
	`

	_, err := d.UserFind(ctx, user.Username)
	if err != nil {
		return fmt.Errorf("%s create userModule: user already exists", userDBErrorStatement)
	}

	if err := d.Client.QueryRow(ctx, stmt, user.Username, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.ID); err != nil {
		return fmt.Errorf("%s create userModule: %s", userDBErrorStatement, err.Error())
	}
	return nil
}
