package racecourse_test

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	handler "keiba-app-backend/handler/race_course"
	model "keiba-app-backend/model/race_course"
)

type mockService struct {
	course model.Racecourse
	err    error
	call   string
}

func (m *mockService) GetRaceCourses() ([]model.Racecourse, error) {
	m.call = "index"
	return []model.Racecourse{m.course}, m.err
}
func (m *mockService) GetRaceCourse(int64) (model.Racecourse, error) {
	m.call = "show"
	return m.course, m.err
}
func (m *mockService) CreateRaceCourse(string, string) (model.Racecourse, error) {
	m.call = "create"
	return m.course, m.err
}
func (m *mockService) UpdateRaceCourse(int64, string, string) (model.Racecourse, error) {
	m.call = "update"
	return m.course, m.err
}
func (m *mockService) DeleteRaceCourse(int64) error { m.call = "delete"; return m.err }

func TestRaceCourseHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected := model.Racecourse{ID: 1, Code: "TKY", Name: "東京"}
	tests := []struct {
		name, method, path, body, call string
		status                         int
		fn                             func(*mockService) gin.HandlerFunc
	}{
		{"index", http.MethodGet, "/race-courses", "", "index", http.StatusOK, func(s *mockService) gin.HandlerFunc { return handler.Index(s) }},
		{"show", http.MethodGet, "/race-courses/1", "", "show", http.StatusOK, func(s *mockService) gin.HandlerFunc { return handler.Show(s) }},
		{"create", http.MethodPost, "/race-courses", `{"code":"TKY","name":"東京"}`, "create", http.StatusCreated, func(s *mockService) gin.HandlerFunc { return handler.Create(s) }},
		{"update", http.MethodPut, "/race-courses/1", `{"code":"TKY","name":"東京"}`, "update", http.StatusOK, func(s *mockService) gin.HandlerFunc { return handler.Update(s) }},
		{"delete", http.MethodDelete, "/race-courses/1", "", "delete", http.StatusNoContent, func(s *mockService) gin.HandlerFunc { return handler.Delete(s) }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &mockService{course: expected}
			c, rec := context(tt.method, tt.path, tt.body)
			tt.fn(service)(c)
			c.Writer.WriteHeaderNow()
			if rec.Code != tt.status || service.call != tt.call {
				t.Fatalf("status=%d call=%s", rec.Code, service.call)
			}
		})
	}
	c, rec := context(http.MethodPost, "/race-courses", `{"code":"","name":"東京"}`)
	handler.Create(&mockService{})(c)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("validation status = %d", rec.Code)
	}
}

func TestRaceCourseHandlersErrors(t *testing.T) {
	for _, fn := range []func(*mockService) gin.HandlerFunc{
		func(s *mockService) gin.HandlerFunc { return handler.Show(s) },
		func(s *mockService) gin.HandlerFunc { return handler.Update(s) },
		func(s *mockService) gin.HandlerFunc { return handler.Delete(s) },
	} {
		c, rec := context(http.MethodGet, "/race-courses/1", `{"code":"TKY","name":"東京"}`)
		fn(&mockService{err: sql.ErrNoRows})(c)
		c.Writer.WriteHeaderNow()
		if rec.Code != http.StatusNotFound {
			t.Errorf("status = %d", rec.Code)
		}
	}
}

func context(method, path, body string) (*gin.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(method, path, bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")
	if len(path) > len("/race-courses/") {
		c.Params = gin.Params{{Key: "id", Value: path[len("/race-courses/"):]}}
	}
	return c, rec
}
