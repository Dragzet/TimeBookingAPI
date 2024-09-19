package bookingModule

import (
	"TimeBookingAPI/internal/storage/PostgreSQL"
	"context"
	"fmt"
)

const errorStatement = "internal/bookingModule/db.go: "

type db struct {
	client PostgreSQL.Client
}

func (d *db) Create(ctx context.Context, booking *BookingModel) error {
	stmt := `
	INSERT INTO bookings
		(username, start_time, end_time)
	VALUES 
	    ($1, $2, $3)
	RETURNING id
	`

	if err := d.client.QueryRow(ctx, stmt, booking.Username, booking.StartTime, booking.EndTime).Scan(&booking.ID); err != nil {
		return fmt.Errorf("%s create bookingModule: %s", errorStatement, err.Error())
	}
	return nil
}

func (d *db) Find(ctx context.Context, id string) (*BookingModel, error) {
	stmt := `
		SELECT id, username, start_time, end_time FROM bookings WHERE username = $1
	`
	var booking BookingModel

	err := d.client.QueryRow(ctx, stmt, id).Scan(&booking.ID, &booking.Username, &booking.StartTime, &booking.EndTime)
	if err != nil {
		return nil, fmt.Errorf("%s find bookingModule: %s", errorStatement, err.Error())
	}
	return &booking, nil
}

func (d *db) FindAll(ctx context.Context, username string) ([]*BookingModel, error) {
	stmt := `
		SELECT id, username, start_time, end_time FROM bookings
	`

	rows, err := d.client.Query(ctx, stmt)
	if err != nil {
		return nil, fmt.Errorf("%s find all bookings: %s", errorStatement, err.Error())
	}

	bokingsArr := make([]*BookingModel, 0)
	for rows.Next() {
		var booking BookingModel
		if err := rows.Scan(&booking.ID, &booking.Username, &booking.StartTime, &booking.EndTime); err != nil {
			return nil, fmt.Errorf("%s find all bookings (reading book): %s", errorStatement, err.Error())
		}
		fmt.Println(booking.Username, username)
		if booking.Username == username {
			bokingsArr = append(bokingsArr, &booking)
		}
	}
	return bokingsArr, nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	stmt := `
		DELETE FROM bookings WHERE id = $1
	`

	if _, err := d.client.Exec(ctx, stmt, id); err != nil {
		return fmt.Errorf("%s delete bookingModule: %s", errorStatement, err.Error())
	}
	return nil
}

func NewDB(client PostgreSQL.Client) BookingStorage {
	return &db{client: client}
}
