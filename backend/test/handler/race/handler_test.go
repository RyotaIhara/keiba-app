package race_test

import (
	"bytes"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	handler "keiba-app-backend/handler/race"
	model "keiba-app-backend/model/race"
	support "keiba-app-backend/model/race/support"
)

type mockRaceService struct {
	race    model.Race
	details []model.RaceDetail
	err     error
	call    string
}

func (m *mockRaceService) GetRace(int64) (model.Race, error) { m.call = "show"; return m.race, m.err }
func (m *mockRaceService) GetRaceDetails(int64) ([]model.RaceDetail, error) {
	m.call = "details"
	return m.details, m.err
}
func (m *mockRaceService) CreateRace(support.RaceInput) (model.Race, error) {
	m.call = "create"
	return m.race, m.err
}
func (m *mockRaceService) UpdateRace(int64, support.RaceInput) (model.Race, error) {
	m.call = "update"
	return m.race, m.err
}
func (m *mockRaceService) DeleteRace(int64) error { m.call = "delete"; return m.err }

type mockSearchService struct {
	races []model.Race
	err   error
	input support.RaceSearchInput
}

func (m *mockSearchService) SearchRaces(input support.RaceSearchInput) ([]model.Race, error) {
	m.input = input
	return m.races, m.err
}

const validRaceJSON = `{"race_date":"2026-09-22","race_course_id":1,"race_number":1,"race_name":"記念","start_time":"12:00:00","surface":"turf","distance":1600,"direction":"right","weather":"sunny","track_condition":"firm","race_conditions":"open"}`

func TestRaceHandlersSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected := model.Race{ID: 1, RaceName: "記念"}
	tests := []struct {
		name, method, path, body, call string
		status                         int
		run                            func(*mockRaceService) gin.HandlerFunc
	}{
		{"show", http.MethodGet, "/races/1", "", "show", http.StatusOK, func(s *mockRaceService) gin.HandlerFunc { return handler.Show(s) }},
		{"details", http.MethodGet, "/races/1/details", "", "details", http.StatusOK, func(s *mockRaceService) gin.HandlerFunc { return handler.DetailsByID(s) }},
		{"create", http.MethodPost, "/races", validRaceJSON, "create", http.StatusCreated, func(s *mockRaceService) gin.HandlerFunc { return handler.Create(s) }},
		{"update", http.MethodPut, "/races/1", validRaceJSON, "update", http.StatusOK, func(s *mockRaceService) gin.HandlerFunc { return handler.Update(s) }},
		{"delete", http.MethodDelete, "/races/1", "", "delete", http.StatusNoContent, func(s *mockRaceService) gin.HandlerFunc { return handler.Delete(s) }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &mockRaceService{race: expected}
			c, rec := context(tt.method, tt.path, tt.body)
			tt.run(service)(c)
			c.Writer.WriteHeaderNow()
			if rec.Code != tt.status || service.call != tt.call {
				t.Fatalf("status=%d call=%s", rec.Code, service.call)
			}
		})
	}
	search := &mockSearchService{races: []model.Race{expected}}
	c, rec := context(http.MethodGet, "/races?race_date=2026-09-22&race_course_id=1", "")
	handler.Index(search)(c)
	if rec.Code != http.StatusOK || search.input.RaceDate == nil || search.input.RaceCourseID == nil {
		t.Fatalf("search response=%d input=%#v", rec.Code, search.input)
	}
}

func TestRaceHandlersValidationAndErrors(t *testing.T) {
	c, rec := context(http.MethodGet, "/races?race_date=bad", "")
	handler.Index(&mockSearchService{})(c)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("invalid query status=%d", rec.Code)
	}
	c, rec = context(http.MethodPost, "/races", `{"race_date":"bad"}`)
	handler.Create(&mockRaceService{})(c)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("invalid request status=%d", rec.Code)
	}
	for _, run := range []func(*mockRaceService) gin.HandlerFunc{
		func(s *mockRaceService) gin.HandlerFunc { return handler.Show(s) },
		func(s *mockRaceService) gin.HandlerFunc { return handler.Update(s) },
		func(s *mockRaceService) gin.HandlerFunc { return handler.Delete(s) },
	} {
		c, rec = context(http.MethodGet, "/races/1", validRaceJSON)
		run(&mockRaceService{err: sql.ErrNoRows})(c)
		if rec.Code != http.StatusNotFound {
			t.Errorf("not found status=%d", rec.Code)
		}
	}
	c, rec = context(http.MethodGet, "/races/1", "")
	serviceErr := errors.New("failed")
	handler.Show(&mockRaceService{err: serviceErr})(c)
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("internal status=%d", rec.Code)
	}
}

func context(method, path, body string) (*gin.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(method, path, bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")
	if len(path) > len("/races/") && path[:len("/races/")] == "/races/" {
		value := path[len("/races/"):]
		if index := bytes.IndexByte([]byte(value), '/'); index >= 0 {
			value = value[:index]
		}
		if index := bytes.IndexByte([]byte(value), '?'); index >= 0 {
			value = value[:index]
		}
		c.Params = gin.Params{{Key: "id", Value: value}}
	}
	return c, rec
}
