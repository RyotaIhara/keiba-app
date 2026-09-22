package racecourse

import (
	"errors"
	"testing"

	raceCourseModel "keiba-app-backend/model/race_course"
	raceCourseService "keiba-app-backend/service/race_course"
)

type mockStore struct {
	courses []raceCourseModel.Racecourse
	course  raceCourseModel.Racecourse
	id      int64
	err     error
}

func (m *mockStore) FetchRaceCourses() ([]raceCourseModel.Racecourse, error) { return m.courses, m.err }
func (m *mockStore) FindRaceCourseByID(int64) (raceCourseModel.Racecourse, error) {
	return m.course, m.err
}
func (m *mockStore) CreateRaceCourse(string, string) (int64, error) { return m.id, m.err }
func (m *mockStore) UpdateRaceCourse(int64, string, string) error   { return m.err }
func (m *mockStore) DeleteRaceCourse(int64) error                   { return m.err }

func TestRaceCourseServiceDelegatesAllMethods(t *testing.T) {
	expected := raceCourseModel.Racecourse{ID: 2, Code: "TKY", Name: "東京"}
	store := &mockStore{courses: []raceCourseModel.Racecourse{expected}, course: expected, id: 2}
	service := raceCourseService.NewRaceCourseService(store)
	if got, err := service.GetRaceCourses(); err != nil || len(got) != 1 || got[0] != expected {
		t.Fatalf("GetRaceCourses = %#v, %v", got, err)
	}
	if got, err := service.GetRaceCourse(2); err != nil || got != expected {
		t.Fatalf("GetRaceCourse = %#v, %v", got, err)
	}
	if got, err := service.CreateRaceCourse("TKY", "東京"); err != nil || got != expected {
		t.Fatalf("CreateRaceCourse = %#v, %v", got, err)
	}
	if got, err := service.UpdateRaceCourse(2, "TKY", "東京"); err != nil || got != expected {
		t.Fatalf("UpdateRaceCourse = %#v, %v", got, err)
	}
	if err := service.DeleteRaceCourse(2); err != nil {
		t.Fatalf("DeleteRaceCourse = %v", err)
	}
}

func TestRaceCourseServicePropagatesErrors(t *testing.T) {
	errExpected := errors.New("store failed")
	store := &mockStore{err: errExpected}
	service := raceCourseService.NewRaceCourseService(store)
	if _, err := service.GetRaceCourses(); !errors.Is(err, errExpected) {
		t.Fatal(err)
	}
	if _, err := service.CreateRaceCourse("a", "b"); !errors.Is(err, errExpected) {
		t.Fatal(err)
	}
	if _, err := service.UpdateRaceCourse(1, "a", "b"); !errors.Is(err, errExpected) {
		t.Fatal(err)
	}
}
