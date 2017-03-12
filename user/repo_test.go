package user

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"fmt"
)

func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreate(t *testing.T) {
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
		UserID: "11",
		Provider: "22",
		Uid: 33,
		Token: "44",
		Ttl: 55,
	}

	r.Create(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestCreateError(t *testing.T) {
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
		UserID: "11",
		Provider: "22",
		Uid: 33,
		Token: "44",
		Ttl: 55,
	}

	r.Create(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestCreateErrorBegin(t *testing.T) {
	mockDb, mock := NewMock(t)
	defer mockDb.Close()

	mock.ExpectBegin().WillReturnError(fmt.Errorf("Генерячим ошибакос"))

	r := userRepo{
		db: mockDb,
	}

	u := &User{
		UserID: "11",
		Provider: "22",
		Uid: 33,
		Token: "44",
		Ttl: 55,
	}

	r.Create(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestCreateErrorRollBack(t *testing.T) {
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
		UserID: "11",
		Provider: "22",
		Uid: 33,
		Token: "44",
		Ttl: 55,
	}

	r.Create(u)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
