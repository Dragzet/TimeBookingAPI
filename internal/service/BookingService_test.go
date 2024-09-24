package service

import (
	"TimeBookingAPI/internal/repository"
	"context"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestService_CreateBooking(t *testing.T) {
	type mockBehavior func(pgxmock.PgxPoolIface, repository.BookingModel, repository.UserModel)

	testTable := []struct {
		name           string
		booking        repository.BookingModel
		user           repository.UserModel
		mockBehavior   mockBehavior
		expectedStatus int
	}{
		{
			name: "OK",
			booking: repository.BookingModel{
				Username:  "admin",
				Delta:     1,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Hour * time.Duration(1)),
			},
			user: repository.UserModel{
				Username: "admin",
				Password: "admins",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, booking repository.BookingModel, user repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(user.Username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO bookings")).
					WithArgs(booking.Username, booking.StartTime, booking.EndTime).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow("123e4567-e89b-12d3-a456-426614174000"))
			},
			expectedStatus: 200,
		},
		{
			name: "BadEndTime",
			booking: repository.BookingModel{
				Username:  "admin",
				Delta:     1,
				StartTime: time.Now(),
				EndTime:   time.Now(),
			},
			user: repository.UserModel{
				Username: "admin",
				Password: "admins",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, booking repository.BookingModel, user repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(user.Username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO bookings")).
					WithArgs(booking.Username, booking.StartTime, booking.EndTime).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow("123e4567-e89b-12d3-a456-426614174000"))
			},
			expectedStatus: 500,
		},
		{
			name: "BadDelta",
			booking: repository.BookingModel{
				Username:  "admin",
				Delta:     -1,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Hour * time.Duration(1)),
			},
			user: repository.UserModel{
				Username: "admin",
				Password: "admins",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, booking repository.BookingModel, user repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(user.Username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO bookings")).
					WithArgs(booking.Username, booking.StartTime, booking.EndTime).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow("123e4567-e89b-12d3-a456-426614174000"))
			},
			expectedStatus: 400,
		},
		{
			name: "BadUser",
			booking: repository.BookingModel{
				Username:  "admin",
				Delta:     1,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Hour * time.Duration(1)),
			},
			user: repository.UserModel{
				Username: "",
				Password: "admins",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, booking repository.BookingModel, user repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(user.Username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO bookings")).
					WithArgs(booking.Username, booking.StartTime, booking.EndTime).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow("123e4567-e89b-12d3-a456-426614174000"))
			},
			expectedStatus: 500,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mock, err := pgxmock.NewPool()
			if err != nil {
				t.Fatal(err)
			}
			defer mock.Close()

			db := repository.NewDB(mock)
			service := NewService(db)

			testCase.mockBehavior(mock, testCase.booking, testCase.user)

			answer := Answer{}
			if answer, err = service.CreateBooking(context.TODO(), &testCase.booking); err != nil {
				if answer.Status != testCase.expectedStatus {
					t.Errorf("unexpected err: %s", err)
				}
			}

			assert.Equal(t, testCase.expectedStatus, answer.Status)
		})
	}

}

func TestService_DeleteBooking(t *testing.T) {
	type mockBehavior func(mock pgxmock.PgxPoolIface, booking repository.BookingModel, target string)

	testTable := []struct {
		name           string
		booking        repository.BookingModel
		targetid       string
		mockBehavior   mockBehavior
		expectedStatus int
	}{
		{
			name: "OK",
			booking: repository.BookingModel{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				Username:  "admin",
				Delta:     1,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Hour * time.Duration(1)),
			},
			targetid: "123e4567-e89b-12d3-a456-426614174000",
			mockBehavior: func(mock pgxmock.PgxPoolIface, booking repository.BookingModel, target string) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, start_time, end_time FROM bookings WHERE id = $1")).
					WithArgs(target).
					WillReturnRows(mock.NewRows([]string{"id", "username", "start_time", "end_time"}).AddRow(booking.ID, booking.Username, booking.StartTime, booking.EndTime))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM bookings WHERE id = $1")).
					WithArgs(target).
					WillReturnResult(pgxmock.NewResult("DELETE", 1))
			},
			expectedStatus: 200,
		},
		{
			name: "BadTargetID",
			booking: repository.BookingModel{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				Username:  "admin",
				Delta:     1,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Hour * time.Duration(1)),
			},
			targetid: "",
			mockBehavior: func(mock pgxmock.PgxPoolIface, booking repository.BookingModel, target string) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, start_time, end_time FROM bookings WHERE id = $1")).
					WithArgs(target).
					WillReturnRows(mock.NewRows([]string{"id", "username", "start_time", "end_time"}).AddRow(booking.ID, booking.Username, booking.StartTime, booking.EndTime))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM bookings WHERE id = $1")).
					WithArgs(target).
					WillReturnResult(pgxmock.NewResult("DELETE", 1))
			},
			expectedStatus: 400,
		},
		{
			name: "BadBookingID",
			booking: repository.BookingModel{
				ID:        "",
				Username:  "admin",
				Delta:     1,
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Hour * time.Duration(1)),
			},
			targetid: "123e4567-e89b-12d3-a456-426614174000",
			mockBehavior: func(mock pgxmock.PgxPoolIface, booking repository.BookingModel, target string) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, start_time, end_time FROM bookings WHERE id = $1")).
					WithArgs(target).
					WillReturnRows(mock.NewRows([]string{"id", "username", "start_time", "end_time"}).AddRow(booking.ID, booking.Username, booking.StartTime, booking.EndTime))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM bookings WHERE id = $1")).
					WithArgs(booking.ID).
					WillReturnResult(pgxmock.NewResult("DELETE", 1))
			},
			expectedStatus: 500,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mock, err := pgxmock.NewPool()
			if err != nil {
				t.Fatal(err)
			}
			defer mock.Close()

			db := repository.NewDB(mock)
			service := NewService(db)

			testCase.mockBehavior(mock, testCase.booking, testCase.targetid)

			answer := Answer{}
			if answer, err = service.DeleteBooking(context.TODO(), testCase.targetid); err != nil {
				if answer.Status != testCase.expectedStatus {
					t.Errorf("unexpected err: %s", err)
				}
			}

			assert.Equal(t, testCase.expectedStatus, answer.Status)
		})
	}
}

func TestService_FindBooking(t *testing.T) {
	type mockBehavior func(mock pgxmock.PgxPoolIface, username string, user repository.UserModel, bookings []*repository.BookingModel)

	testTable := []struct {
		name           string
		username       string
		user           repository.UserModel
		bookings       []*repository.BookingModel
		mockBehavior   mockBehavior
		expectedStatus int
	}{
		{
			name:     "OK",
			username: "admin",
			user: repository.UserModel{
				Username: "admin",
				Password: "admin",
			},
			bookings: []*repository.BookingModel{
				{
					ID:        "123e4567-e89b-12d3-a456-426614174000",
					Username:  "admin",
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour * time.Duration(1)),
					Delta:     1,
				},
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, username string, user repository.UserModel, bookings []*repository.BookingModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				rows := mock.NewRows([]string{"id", "username", "start_time", "end_time"})
				for _, booking := range bookings {
					rows.AddRow(booking.ID, booking.Username, booking.StartTime, booking.EndTime)
				}

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, start_time, end_time FROM bookings")).
					WillReturnRows(rows)
			},
			expectedStatus: 200,
		},
		{
			name:     "EmptyUsername",
			username: "",
			user: repository.UserModel{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				Username:  "admin",
				Password:  "admin",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			bookings: []*repository.BookingModel{
				{
					ID:        "123e4567-e89b-12d3-a456-426614174000",
					Username:  "admin",
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour * time.Duration(1)),
					Delta:     1,
				},
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, username string, user repository.UserModel, bookings []*repository.BookingModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow(user.ID, user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				rows := mock.NewRows([]string{"id", "username", "start_time", "end_time"})
				for _, booking := range bookings {
					rows.AddRow(booking.ID, booking.Username, booking.StartTime, booking.EndTime)
				}

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, start_time, end_time FROM bookings")).
					WillReturnRows(rows)
			},
			expectedStatus: 400,
		},
		{
			name:     "EmptyUsername",
			username: "admin",
			user: repository.UserModel{
				ID:        "123e4567-e89b-12d3-a456-426614174000",
				Username:  "",
				Password:  "admin",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			bookings: []*repository.BookingModel{
				{
					ID:        "123e4567-e89b-12d3-a456-426614174000",
					Username:  "admin",
					StartTime: time.Now(),
					EndTime:   time.Now().Add(time.Hour * time.Duration(1)),
					Delta:     1,
				},
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, username string, user repository.UserModel, bookings []*repository.BookingModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs("").
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow(user.ID, user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				rows := mock.NewRows([]string{"id", "username", "start_time", "end_time"})
				for _, booking := range bookings {
					rows.AddRow(booking.ID, booking.Username, booking.StartTime, booking.EndTime)
				}

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, start_time, end_time FROM bookings")).
					WillReturnRows(rows)
			},
			expectedStatus: 500,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mock, err := pgxmock.NewPool()
			if err != nil {
				t.Fatal(err)
			}
			defer mock.Close()

			db := repository.NewDB(mock)
			service := NewService(db)

			testCase.mockBehavior(mock, testCase.username, testCase.user, testCase.bookings)

			answer := Answer{}
			if answer, err = service.FindBooking(context.TODO(), testCase.username); err != nil {
				if answer.Status != testCase.expectedStatus {
					t.Errorf("unexpected err: %s", err)
				}
			}

			assert.Equal(t, testCase.expectedStatus, answer.Status)
		})
	}
}
