package user

import (
	"database/sql"
	"testing"

	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestUser_Create(t *testing.T) {
	mockDb, mock := NewMock(t)
	defer mockDb.Close()

	mock.ExpectBegin()
	ee := mock.ExpectExec("INSERT")
	ee.WithArgs(
		"11",
		"22",
		33,
		"44",
		55,
	)
	ee.WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := userRepo{
		db: mockDb,
	}

	u := &User{
		UserID:   "11",
		Provider: "22",
		Uid:      33,
		Token:    "44",
		Ttl:      55,
	}

	r.Create(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUser_Create_Error(t *testing.T) {
	mockDb, mock := NewMock(t)
	defer mockDb.Close()

	mock.ExpectBegin()
	ee := mock.ExpectExec("INSERT")
	ee.WithArgs(
		"11",
		"22",
		33,
		"44",
		55,
	)
	ee.WillReturnError(fmt.Errorf("Генерячим ошибакос"))
	mock.ExpectRollback()

	r := userRepo{
		db: mockDb,
	}

	u := &User{
		UserID:   "11",
		Provider: "22",
		Uid:      33,
		Token:    "44",
		Ttl:      55,
	}

	r.Create(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUser_Create_Error_Begin(t *testing.T) {
	mockDb, mock := NewMock(t)
	defer mockDb.Close()

	mock.ExpectBegin().WillReturnError(fmt.Errorf("Генерячим ошибакос"))

	r := userRepo{
		db: mockDb,
	}

	u := &User{
		UserID:   "11",
		Provider: "22",
		Uid:      33,
		Token:    "44",
		Ttl:      55,
	}

	r.Create(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUser_Create_Error_RollBack(t *testing.T) {
	mockDb, mock := NewMock(t)
	defer mockDb.Close()

	mock.ExpectBegin()
	ee := mock.ExpectExec("INSERT")
	ee.WithArgs(
		"11",
		"22",
		33,
		"44",
		55,
	)
	ee.WillReturnError(fmt.Errorf("Генерячим ошибакос"))
	mock.ExpectRollback().WillReturnError(fmt.Errorf("Генерячим ошибакос"))

	r := userRepo{
		db: mockDb,
	}

	u := &User{
		UserID:   "11",
		Provider: "22",
		Uid:      33,
		Token:    "44",
		Ttl:      55,
	}

	r.Create(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUser_Update(t *testing.T) {
	mockDb, mock := NewMock(t)
	defer mockDb.Close()

	mock.ExpectBegin()
	ee := mock.ExpectExec("UPDATE")
	ee.WithArgs(
		"1",
		2,
		"11",
	)
	ee.WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := userRepo{
		db: mockDb,
	}

	u := &User{
		UserID: "11",
		Token:  "1",
		Ttl:    2,
	}

	r.Update(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUser_Update_Error(t *testing.T) {
	mockDb, mock := NewMock(t)
	defer mockDb.Close()

	mock.ExpectBegin()
	ee := mock.ExpectExec("UPDATE")
	ee.WithArgs(
		"1",
		2,
		"11",
	)
	ee.WillReturnError(fmt.Errorf("Gen error"))
	mock.ExpectRollback()

	r := userRepo{
		db: mockDb,
	}

	u := &User{
		UserID: "11",
		Token:  "1",
		Ttl:    2,
	}

	r.Update(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
