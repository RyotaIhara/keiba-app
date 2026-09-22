package race

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	raceInfrastructure "keiba-app-backend/infrastructure/race"
	raceSupport "keiba-app-backend/model/race/support"
	raceTypes "keiba-app-backend/model/race/types"
)

func newRaceStore(t *testing.T) (*raceInfrastructure.Store, sqlmock.Sqlmock, func()) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	return raceInfrastructure.NewStore(db), mock, func() { db.Close() }
}

func raceColumns() []string {
	return []string{"id", "race_date", "course_id", "course_code", "course_name", "race_number", "race_name", "start_time", "surface", "distance", "direction", "weather", "track_condition", "race_conditions"}
}
func raceRows() *sqlmock.Rows {
	return sqlmock.NewRows(raceColumns()).AddRow(1, time.Date(2026, 9, 22, 0, 0, 0, 0, time.UTC), 2, "TKY", "東京", 1, "記念", "12:00:00", "turf", 1600, "right", "sunny", "firm", "open")
}
func detailRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"detail_id", "id", "race_date", "course_id", "course_code", "course_name", "race_number", "race_name", "start_time", "surface", "distance", "direction", "weather", "track_condition", "race_conditions", "horse_number", "frame_number", "horse_name", "sex", "age", "weight", "jockey", "stable", "body_weight", "body_weight_change", "odds", "popularity"}).
		AddRow(10, 1, time.Date(2026, 9, 22, 0, 0, 0, 0, time.UTC), 2, "TKY", "東京", 1, "記念", "12:00:00", "turf", 1600, "right", "sunny", "firm", "open", 1, 1, "Horse", "colt", 3, 56.0, "Jockey", "Stable", 480, 2, 3.2, 1)
}
func input() raceSupport.RaceInput {
	return raceSupport.RaceInput{RaceDate: time.Date(2026, 9, 22, 0, 0, 0, 0, time.UTC), RaceCourseID: 2, RaceNumber: 1, RaceName: "記念", StartTime: time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC), Surface: raceTypes.SurfaceTurf, Distance: 1600, Direction: raceTypes.DirectionRight, Weather: raceTypes.WeatherSunny, TrackCondition: raceTypes.TrackConditionFirm, RaceConditions: "open"}
}

func TestStoreRaceQueries(t *testing.T) {
	store, mock, closeDB := newRaceStore(t)
	defer closeDB()
	mock.ExpectQuery("SELECT").WillReturnRows(raceRows())
	if got, err := store.FetchRaces(); err != nil || got[0].Racecourse.Name != "東京" {
		t.Fatalf("FetchRaces = %#v, %v", got, err)
	}
	mock.ExpectQuery("SELECT").WithArgs(time.Date(2026, 9, 22, 0, 0, 0, 0, time.UTC), int64(2)).WillReturnRows(raceRows())
	date := time.Date(2026, 9, 22, 0, 0, 0, 0, time.UTC)
	courseID := int64(2)
	if got, err := store.SearchRaces(raceSupport.RaceSearchInput{RaceDate: &date, RaceCourseID: &courseID}); err != nil || len(got) != 1 {
		t.Fatalf("SearchRaces = %#v, %v", got, err)
	}
	mock.ExpectQuery("SELECT").WithArgs(int64(1)).WillReturnRows(raceRows())
	if got, err := store.FindRaceByID(1); err != nil || got.ID != 1 {
		t.Fatalf("FindRaceByID = %#v, %v", got, err)
	}
	mock.ExpectQuery("SELECT").WithArgs(int64(1)).WillReturnRows(detailRows())
	if got, err := store.FetchRaceDetailsByRaceID(1); err != nil || got[0].Race.ID != 1 {
		t.Fatalf("FetchRaceDetailsByRaceID = %#v, %v", got, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestStoreRaceMutations(t *testing.T) {
	store, mock, closeDB := newRaceStore(t)
	defer closeDB()
	in := input()
	mock.ExpectExec("INSERT INTO races").WithArgs(in.RaceDate, in.RaceCourseID, in.RaceNumber, in.RaceName, in.StartTime, in.Surface, in.Distance, in.Direction, in.Weather, in.TrackCondition, in.RaceConditions).WillReturnResult(sqlmock.NewResult(5, 1))
	if id, err := store.CreateRace(in); err != nil || id != 5 {
		t.Fatalf("CreateRace = %d, %v", id, err)
	}
	mock.ExpectExec("UPDATE races").WillReturnResult(sqlmock.NewResult(0, 1))
	if err := store.UpdateRace(5, in); err != nil {
		t.Fatal(err)
	}
	mock.ExpectExec("DELETE FROM races").WithArgs(int64(5)).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := store.DeleteRace(5); err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestStoreRaceErrorsAndNotFound(t *testing.T) {
	store, mock, closeDB := newRaceStore(t)
	defer closeDB()
	dbErr := errors.New("db failed")
	mock.ExpectQuery("SELECT").WillReturnError(dbErr)
	if _, err := store.FetchRaces(); !errors.Is(err, dbErr) {
		t.Fatal(err)
	}
	mock.ExpectExec("DELETE FROM races").WithArgs(int64(9)).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT").WithArgs(int64(9)).WillReturnError(sql.ErrNoRows)
	if err := store.DeleteRace(9); !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("DeleteRace = %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
