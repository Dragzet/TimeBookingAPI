package userModule

import (
	"TimeBookingAPI/internal/bookingModule"
	"TimeBookingAPI/internal/storage/PostgreSQL"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const errorStatement = "internal/userModule/db.go: "

type db struct {
	client PostgreSQL.Client
}

func (d db) Create(ctx context.Context, user *UserModel) error {
	stmt := `
	INSERT INTO users
		(username, password, created_at, updated_at)
	VALUES 
	    ($1, $2, $3, $4)
	RETURNING id
	`

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s create hashing userModule: %s", errorStatement, err.Error())
	}

	if err := d.client.QueryRow(ctx, stmt, user.Username, hash, user.CreatedAt, user.UpdatedAt).Scan(&user.ID); err != nil {
		return fmt.Errorf("%s create userModule: %s", errorStatement, err.Error())
	}
	return nil
}

func (d db) Find(ctx context.Context, username string) (*UserModel, error) {
	stmt := `
		SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1
	`
	var user UserModel

	err := d.client.QueryRow(ctx, stmt, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s find userModule: %s", errorStatement, err.Error())
	}
	return &user, nil
}

func (d db) Update(ctx context.Context, item *UserModel) error {
	stmt := `
	UPDATE users SET username = $2, password = $3 WHERE id = $1;
	`

	user, err := d.Find(ctx, item.Username)
	if err != nil {
		return fmt.Errorf("%s delete userModule findUser: %s", errorStatement, err.Error())
	}

	item.ID = user.ID

	if _, err := d.client.Exec(ctx, stmt, item.ID, item.Username, item.Password); err != nil {
		return fmt.Errorf("%s update userModule: %s", errorStatement, err.Error())
	}
	return nil
}

func (d db) Delete(ctx context.Context, bookings bookingModule.BookingStorage, username string) error {
	stmt := `
		DELETE FROM users WHERE id = $1
	`

	user, err := d.Find(ctx, username)
	if err != nil {
		return fmt.Errorf("%s delete userModule findUser: %s", errorStatement, err.Error())
	}

	bookingsArr, err := bookings.FindAll(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("%s delete userModule findAll: %s", errorStatement, err.Error())
	}

	for _, tempBooking := range bookingsArr {
		if err := bookings.Delete(ctx, tempBooking.ID); err != nil {
			return fmt.Errorf("%s delete userModule DeleteBooking: %s", errorStatement, err.Error())
		}
	}

	if _, err := d.client.Exec(ctx, stmt, user.ID); err != nil {
		return fmt.Errorf("%s delete userModule: %s", errorStatement, err.Error())
	}

	return nil
}

func NewDB(client PostgreSQL.Client) UserStorage {
	return &db{client: client}
}
