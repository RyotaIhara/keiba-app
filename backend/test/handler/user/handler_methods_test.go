package user_test

import (
	"bytes"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	userHandler "keiba-app-backend/handler/user"
	userModel "keiba-app-backend/model/user"
)

type configurableUserService struct {
	user  userModel.User
	err   error
	calls []string
}

func (m *configurableUserService) GetUsers() ([]userModel.User, error) {
	m.calls = append(m.calls, "list")
	return []userModel.User{m.user}, m.err
}
func (m *configurableUserService) GetUser(int64) (userModel.User, error) {
	m.calls = append(m.calls, "show")
	return m.user, m.err
}
func (m *configurableUserService) CreateUser(string, string, string) (userModel.User, error) {
	m.calls = append(m.calls, "create")
	return m.user, m.err
}
func (m *configurableUserService) UpdateUser(int64, string, string) (userModel.User, error) {
	m.calls = append(m.calls, "update")
	return m.user, m.err
}
func (m *configurableUserService) DeleteUser(int64) error {
	m.calls = append(m.calls, "delete")
	return m.err
}

func TestUserHandlersSuccessAndValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected := userModel.User{ID: 1, Code: "u1", Name: "User"}
	tests := []struct {
		name               string
		handler            gin.HandlerFunc
		method, path, body string
		status             int
		call               string
	}{
		{"show", userHandler.Show(&configurableUserService{user: expected}), http.MethodGet, "/users/1", "", http.StatusOK, "show"},
		{"create", userHandler.Create(&configurableUserService{user: expected}), http.MethodPost, "/users", `{"code":"u1","name":"User","password":"pw"}`, http.StatusCreated, "create"},
		{"update", userHandler.Update(&configurableUserService{user: expected}), http.MethodPut, "/users/1", `{"code":"u1","name":"User"}`, http.StatusOK, "update"},
		{"delete", userHandler.Delete(&configurableUserService{}), http.MethodDelete, "/users/1", "", http.StatusNoContent, "delete"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &configurableUserService{user: expected}
			handler := map[string]gin.HandlerFunc{
				"show": userHandler.Show(service), "create": userHandler.Create(service),
				"update": userHandler.Update(service), "delete": userHandler.Delete(service),
			}[tt.call]
			c, rec := userContext(tt.method, tt.path, tt.body)
			handler(c)
			c.Writer.WriteHeaderNow()
			if rec.Code != tt.status {
				t.Fatalf("status = %d, want %d", rec.Code, tt.status)
			}
			if len(service.calls) != 1 || service.calls[0] != tt.call {
				t.Fatalf("calls = %#v", service.calls)
			}
		})
	}
	c, rec := userContext(http.MethodGet, "/users/not-an-id", "")
	userHandler.Show(&configurableUserService{})(c)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("invalid ID status = %d", rec.Code)
	}
	c, rec = userContext(http.MethodPost, "/users", `{"code":"","name":"User","password":"pw"}`)
	userHandler.Create(&configurableUserService{})(c)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("invalid body status = %d", rec.Code)
	}
}

func TestUserHandlersMapNotFoundAndInternalErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)
	for _, handler := range []gin.HandlerFunc{
		userHandler.Show(&configurableUserService{err: sql.ErrNoRows}),
		userHandler.Update(&configurableUserService{err: sql.ErrNoRows}),
		userHandler.Delete(&configurableUserService{err: sql.ErrNoRows}),
	} {
		c, rec := userContext(http.MethodGet, "/users/1", `{"code":"u","name":"n"}`)
		handler(c)
		if rec.Code != http.StatusNotFound {
			t.Errorf("not found status = %d", rec.Code)
		}
	}
	c, rec := userContext(http.MethodGet, "/users", "")
	serviceErr := errors.New("failed")
	userHandler.Index(&configurableUserService{err: serviceErr})(c)
	if rec.Code != http.StatusInternalServerError || len(c.Errors) != 1 {
		t.Fatalf("index error response: %d, errors=%d", rec.Code, len(c.Errors))
	}
}

func userContext(method, path, body string) (*gin.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(method, path, bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")
	if path == "/users/1" || path == "/users/not-an-id" {
		value := path[len("/users/"):]
		c.Params = gin.Params{{Key: "id", Value: value}}
	}
	return c, rec
}
