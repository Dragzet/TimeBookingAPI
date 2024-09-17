package booking

import (
	"TimeBookingAPI/iternal/storage/PostgreSQL"
	"context"
	"fmt"
)

const errorStatement = "iternal/booking/db.go: "

type db struct {
	client PostgreSQL.Client
}

func (d *db) Create(ctx context.Context, item *Booking) error {
	stmt := `
	INSERT INTO bookings
		(user_id, start_time, end_time)
	VALUES 
	    ($1, $2, $3)
	RETURNING id
	`
	if err := d.client.QueryRow(ctx, stmt, item.UserID, item.StartTime, item.EndTime).Scan(&item.ID); err != nil {
		return fmt.Errorf("%s create booking: %s", errorStatement, err.Error())
	}
	return nil
}

func (d *db) Find(ctx context.Context, id string) (*Booking, error) {
	stmt := `
		SELECT id, user_id, start_time, end_time FROM bookings WHERE id = $1
	`

	var booking Booking
	err := d.client.QueryRow(ctx, stmt, id).Scan(&booking.ID, &booking.UserID, &booking.StartTime, &booking.EndTime)
	if err != nil {
		return nil, fmt.Errorf("%s find booking: %s", errorStatement, err.Error())
	}
	return &booking, nil
}

func (d *db) FindAll(ctx context.Context, userID string) ([]*Booking, error) {
	stmt := `
		SELECT id, user_id, start_time, end_time FROM bookings
	`

	rows, err := d.client.Query(ctx, stmt)
	if err != nil {
		return nil, fmt.Errorf("%s find all bookings: %s", errorStatement, err.Error())
	}

	bokingsArr := make([]*Booking, 0)
	for rows.Next() {
		var booking Booking
		if err := rows.Scan(&booking.ID, &booking.UserID, &booking.StartTime, &booking.EndTime); err != nil {
			return nil, fmt.Errorf("%s find all bookings (reading book): %s", errorStatement, err.Error())
		}
		if booking.UserID == userID {
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
		return fmt.Errorf("%s delete booking: %s", errorStatement, err.Error())
	}
	return nil
}

func NewDB(client PostgreSQL.Client) BookingStorage {
	return &db{client: client}
}
