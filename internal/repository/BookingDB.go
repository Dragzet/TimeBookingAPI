package repository

import (
	"context"
	"fmt"
)

const bookingDBErrorStatement = "internal/bookingModule/UserDB.go: "

func (d *DB) BookingCreate(ctx context.Context, booking *BookingModel) error {
	stmt := `
	INSERT INTO bookings
		(username, start_time, end_time)
	VALUES 
	    ($1, $2, $3)
	RETURNING id
	`

	_, err := d.UserFind(ctx, booking.Username)
	if err != nil {
		return fmt.Errorf("%s create bookingModule: %s", bookingDBErrorStatement, err.Error())
	}

	if err := d.Client.QueryRow(ctx, stmt, booking.Username, booking.StartTime, booking.EndTime).Scan(&booking.ID); err != nil {
		return fmt.Errorf("%s create bookingModule: %s", bookingDBErrorStatement, err.Error())
	}
	return nil
}

func (d *DB) BookingFindAll(ctx context.Context, username string) ([]*BookingModel, error) {
	stmt := `
		SELECT id, username, start_time, end_time FROM bookings
	`

	_, err := d.UserFind(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("%s find all bookings: %s", bookingDBErrorStatement, err.Error())
	}

	rows, err := d.Client.Query(ctx, stmt)
	if err != nil {
		return nil, fmt.Errorf("%s find all bookings: %s", bookingDBErrorStatement, err.Error())
	}

	bokingsArr := make([]*BookingModel, 0)
	for rows.Next() {
		var booking BookingModel
		if err := rows.Scan(&booking.ID, &booking.Username, &booking.StartTime, &booking.EndTime); err != nil {
			return nil, fmt.Errorf("%s find all bookings (reading book): %s", bookingDBErrorStatement, err.Error())
		}
		if booking.Username == username {
			bokingsArr = append(bokingsArr, &booking)
		}
	}
	return bokingsArr, nil
}

func (d *DB) BookingDelete(ctx context.Context, id string) error {
	stmt := `
		DELETE FROM bookings WHERE id = $1
	`

	_, err := d.BookingFind(ctx, id)
	if err != nil {
		return fmt.Errorf("%s delete bookingModule: %s", bookingDBErrorStatement, err.Error())
	}

	if _, err := d.Client.Exec(ctx, stmt, id); err != nil {
		return fmt.Errorf("%s delete bookingModule: %s", bookingDBErrorStatement, err.Error())
	}
	return nil
}

func (d *DB) BookingFind(ctx context.Context, id string) (*BookingModel, error) {
	stmt := `
		SELECT id, username, start_time, end_time FROM bookings WHERE id = $1
	`
	var booking BookingModel

	err := d.Client.QueryRow(ctx, stmt, id).Scan(&booking.ID, &booking.Username, &booking.StartTime, &booking.EndTime)
	if err != nil {
		return nil, fmt.Errorf("%s find bookingModule: %s", bookingDBErrorStatement, err.Error())
	}
	return &booking, nil
}
