package user

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	userInfrastructure "keiba-app-backend/infrastructure/user"
)

func newUserStore(t *testing.T) (*userInfrastructure.Store, sqlmock.Sqlmock, func()) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	return userInfrastructure.NewStore(db), mock, func() { db.Close() }
}

func userRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "code", "name", "password"}).AddRow(1, "u1", "User", "pw")
}

func TestStoreUserQueries(t *testing.T) {
	store, mock, closeDB := newUserStore(t)
	defer closeDB()
	mock.ExpectQuery("SELECT id, code, name, password").WillReturnRows(userRows())
	users, err := store.FetchUsers()
	if err != nil || len(users) != 1 || users[0].ID != 1 {
		t.Fatalf("FetchUsers = %#v, %v", users, err)
	}
	mock.ExpectQuery("SELECT id, code, name, password").WithArgs("u1", "pw").WillReturnRows(userRows())
	if got, err := store.FindUserByCodeAndPass("u1", "pw"); err != nil || got.Code != "u1" {
		t.Fatalf("FindUserByCodeAndPass = %#v, %v", got, err)
	}
	mock.ExpectQuery("SELECT id, code, name, password").WithArgs(int64(1)).WillReturnRows(userRows())
	if got, err := store.FindUserByID(1); err != nil || got.ID != 1 {
		t.Fatalf("FindUserByID = %#v, %v", got, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestStoreUserMutations(t *testing.T) {
	store, mock, closeDB := newUserStore(t)
	defer closeDB()
	mock.ExpectExec("INSERT INTO users").WithArgs("u", "User", "pw").WillReturnResult(sqlmock.NewResult(5, 1))
	if id, err := store.CreateUser("u", "User", "pw"); err != nil || id != 5 {
		t.Fatalf("CreateUser = %d, %v", id, err)
	}
	mock.ExpectExec("UPDATE users").WithArgs("u2", "User 2", int64(5)).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := store.UpdateUser(5, "u2", "User 2"); err != nil {
		t.Fatal(err)
	}
	mock.ExpectExec("DELETE FROM users").WithArgs(int64(5)).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := store.DeleteUser(5); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestStoreUserPropagatesQueryAndNotFoundErrors(t *testing.T) {
	store, mock, closeDB := newUserStore(t)
	defer closeDB()
	dbErr := errors.New("db failed")
	mock.ExpectQuery("SELECT id, code, name, password").WillReturnError(dbErr)
	if _, err := store.FetchUsers(); !errors.Is(err, dbErr) {
		t.Fatal(err)
	}
	mock.ExpectExec("DELETE FROM users").WithArgs(int64(9)).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT id, code, name, password").WithArgs(int64(9)).WillReturnError(sql.ErrNoRows)
	if err := store.DeleteUser(9); !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("DeleteUser = %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
