package user_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	userHandler "keiba-app-backend/handler/user"
	userModel "keiba-app-backend/model/user"
)

type mockUserService struct {
	users []userModel.User
	err   error
}

func (m mockUserService) GetUsers() ([]userModel.User, error) {
	return m.users, m.err
}

func newTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(http.MethodGet, "/api/users", nil)
	return c, rec
}

func TestIndexReturnsUsersAsJSON(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	c, rec := newTestContext()
	expected := []userModel.User{
		{ID: 1, Code: "test001", Name: "テスト001", Password: "password"},
		{ID: 2, Code: "test002", Name: "テスト002", Password: "password"},
	}

	// Act
	userHandler.Index(mockUserService{users: expected})(c)

	// Assert
	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}
	if got := rec.Body.String(); got == "" {
		t.Fatal("expected a JSON response body")
	}

	var actual []userModel.User
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}
	if len(actual) != len(expected) {
		t.Fatalf("expected %d users, got %d", len(expected), len(actual))
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("expected user %d to be %+v, got %+v", i, expected[i], actual[i])
		}
	}
}

func TestIndexReturnsInternalServerErrorWhenServiceFails(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	c, rec := newTestContext()
	serviceErr := errors.New("database unavailable")

	// Act
	userHandler.Index(mockUserService{err: serviceErr})(c)

	// Assert
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, rec.Code)
	}
	if got := rec.Body.String(); got != "{\"error\":\"failed to fetch users\"}" {
		t.Fatalf("expected error response, got %q", got)
	}
	if len(c.Errors) != 1 || !errors.Is(c.Errors[0].Err, serviceErr) {
		t.Fatalf("expected service error to be recorded in Gin context")
	}
}

var _ interface {
	GetUsers() ([]userModel.User, error)
} = mockUserService{}
