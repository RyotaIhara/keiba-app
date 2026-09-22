package racecourse

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	raceCourseInfrastructure "keiba-app-backend/infrastructure/race_course"
)

func newRaceCourseStore(t *testing.T) (*raceCourseInfrastructure.Store, sqlmock.Sqlmock, func()) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	return raceCourseInfrastructure.NewStore(db), mock, func() { db.Close() }
}

func courseRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "code", "name"}).AddRow(1, "TKY", "東京")
}

func TestStoreRaceCourseQueriesAndMutations(t *testing.T) {
	store, mock, closeDB := newRaceCourseStore(t)
	defer closeDB()
	mock.ExpectQuery("SELECT id, code, name").WillReturnRows(courseRows())
	if got, err := store.FetchRaceCourses(); err != nil || got[0].Name != "東京" {
		t.Fatalf("FetchRaceCourses = %#v, %v", got, err)
	}
	mock.ExpectQuery("SELECT id, code, name").WithArgs(int64(1)).WillReturnRows(courseRows())
	if got, err := store.FindRaceCourseByID(1); err != nil || got.ID != 1 {
		t.Fatalf("FindRaceCourseByID = %#v, %v", got, err)
	}
	mock.ExpectExec("INSERT INTO race_courses").WithArgs("TKY", "東京").WillReturnResult(sqlmock.NewResult(2, 1))
	if id, err := store.CreateRaceCourse("TKY", "東京"); err != nil || id != 2 {
		t.Fatalf("CreateRaceCourse = %d, %v", id, err)
	}
	mock.ExpectExec("UPDATE race_courses").WithArgs("TKY2", "東京2", int64(1)).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := store.UpdateRaceCourse(1, "TKY2", "東京2"); err != nil {
		t.Fatal(err)
	}
	mock.ExpectExec("DELETE FROM race_courses").WithArgs(int64(1)).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := store.DeleteRaceCourse(1); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestStoreRaceCourseErrors(t *testing.T) {
	store, mock, closeDB := newRaceCourseStore(t)
	defer closeDB()
	dbErr := errors.New("db failed")
	mock.ExpectQuery("SELECT id, code, name").WillReturnError(dbErr)
	if _, err := store.FetchRaceCourses(); !errors.Is(err, dbErr) {
		t.Fatal(err)
	}
	mock.ExpectExec("DELETE FROM race_courses").WithArgs(int64(4)).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT id, code, name").WithArgs(int64(4)).WillReturnError(sql.ErrNoRows)
	if err := store.DeleteRaceCourse(4); !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("DeleteRaceCourse = %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
