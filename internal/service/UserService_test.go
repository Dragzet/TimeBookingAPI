package service

import (
	"TimeBookingAPI/internal/repository"
	"context"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestService_CreateUser(t *testing.T) {
	type mockBehavior func(mock pgxmock.PgxPoolIface, user *repository.UserModel)

	testTable := []struct {
		name           string
		user           *repository.UserModel
		mockBehavior   mockBehavior
		expectedStatus int
	}{
		{
			name: "OK",
			user: &repository.UserModel{
				Username: "admin",
				Password: "admin",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, user *repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(user.Username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4)")).
					WithArgs(user.Username, user.Password, user.CreatedAt, user.UpdatedAt).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow("123e4567-e89b-12d3-a456-426614174100"))
			},
			expectedStatus: 200,
		},
		{
			name: "BadUsername",
			user: &repository.UserModel{
				Username: "",
				Password: "admin",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, user *repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(user.Username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4)")).
					WithArgs(user.Username, user.Password, user.CreatedAt, user.UpdatedAt).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow("123e4567-e89b-12d3-a456-426614174100"))
			},
			expectedStatus: 400,
		},
		{
			name: "BadPassword",
			user: &repository.UserModel{
				Username: "admin",
				Password: "1",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, user *repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(user.Username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4)")).
					WithArgs(user.Username, user.Password, user.CreatedAt, user.UpdatedAt).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow("123e4567-e89b-12d3-a456-426614174100"))
			},
			expectedStatus: 400,
		},
		{
			name: "AlreadyExists",
			user: &repository.UserModel{
				Username: "admin",
				Password: "11231",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, user *repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs("").
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4)")).
					WithArgs(user.Username, user.Password, user.CreatedAt, user.UpdatedAt).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow(""))
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

			testCase.mockBehavior(mock, testCase.user)

			answer := Answer{}
			if answer, err = service.TestUserCreate(context.TODO(), testCase.user); err != nil {
				if answer.Status != testCase.expectedStatus {
					t.Errorf("unexpected err: %s", err)
				}
			}

			assert.Equal(t, testCase.expectedStatus, answer.Status)
		})
	}
}

func TestService_DeleteUser(t *testing.T) {
	type mockBehavior func(mock pgxmock.PgxPoolIface, username string, user repository.UserModel)

	testTable := []struct {
		name           string
		username       string
		user           repository.UserModel
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
			mockBehavior: func(mock pgxmock.PgxPoolIface, username string, user repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				rows := mock.NewRows([]string{"id", "username", "start_time", "end_time"})
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, start_time, end_time FROM bookings")).
					WillReturnRows(rows)

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnResult(pgxmock.NewResult("DELETE", 1))
			},
			expectedStatus: 200,
		},
		{
			name:     "BadUsername",
			username: "",
			user: repository.UserModel{
				Username: "",
				Password: "admin",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, username string, user repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				rows := mock.NewRows([]string{"id", "username", "start_time", "end_time"})
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, start_time, end_time FROM bookings")).
					WillReturnRows(rows)

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnResult(pgxmock.NewResult("DELETE", 1))
			},
			expectedStatus: 400,
		},
		{
			name:     "NoUserFound",
			username: "asd",
			user: repository.UserModel{
				Username: "",
				Password: "admin",
			},
			mockBehavior: func(mock pgxmock.PgxPoolIface, username string, user repository.UserModel) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1")).
					WithArgs(username).
					WillReturnRows(mock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow("123e4567-e89b-12d3-a456-426614174000", user.Username, user.Password, user.CreatedAt, user.UpdatedAt))

				rows := mock.NewRows([]string{"id", "username", "start_time", "end_time"})
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, start_time, end_time FROM bookings")).
					WillReturnRows(rows)

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE username = $1")).
					WithArgs(username).
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

			testCase.mockBehavior(mock, testCase.username, testCase.user)

			answer := Answer{}
			if answer, err = service.DeleteUser(context.TODO(), testCase.username); err != nil {
				if answer.Status != testCase.expectedStatus {
					t.Errorf("unexpected err: %s", err)
				}
			}

			assert.Equal(t, testCase.expectedStatus, answer.Status)
		})
	}
}
