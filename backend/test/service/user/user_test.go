package user

import (
	"errors"
	"testing"

	userModel "keiba-app-backend/model/user"
	userService "keiba-app-backend/service/user"
)

type mockStore struct {
	users      []userModel.User
	user       userModel.User
	id         int64
	err        error
	fetchCalls int
	findCalls  []int64
	createArgs []string
	updateArgs []any
	deleteID   int64
}

func (m *mockStore) FetchUsers() ([]userModel.User, error) { m.fetchCalls++; return m.users, m.err }
func (m *mockStore) FindUserByID(id int64) (userModel.User, error) {
	m.findCalls = append(m.findCalls, id)
	return m.user, m.err
}
func (m *mockStore) CreateUser(code, name, password string) (int64, error) {
	m.createArgs = []string{code, name, password}
	return m.id, m.err
}
func (m *mockStore) UpdateUser(id int64, code, name string) error {
	m.updateArgs = []any{id, code, name}
	return m.err
}
func (m *mockStore) DeleteUser(id int64) error { m.deleteID = id; return m.err }

func TestServiceDelegatesReadAndDeleteMethods(t *testing.T) {
	expected := userModel.User{ID: 1, Code: "u1"}
	store := &mockStore{users: []userModel.User{expected}, user: expected}
	service := userService.NewService(store)

	users, err := service.GetUsers()
	if err != nil || len(users) != 1 || users[0] != expected {
		t.Fatalf("GetUsers = %#v, %v", users, err)
	}
	got, err := service.GetUser(1)
	if err != nil || got != expected || store.findCalls[0] != 1 {
		t.Fatalf("GetUser = %#v, %v", got, err)
	}
	if err := service.DeleteUser(2); err != nil || store.deleteID != 2 {
		t.Fatalf("DeleteUser = %v", err)
	}
}

func TestServiceCreateAndUpdateReturnStoredUser(t *testing.T) {
	expected := userModel.User{ID: 7, Code: "updated"}
	store := &mockStore{id: 7, user: expected}
	service := userService.NewService(store)

	got, err := service.CreateUser("u", "User", "pw")
	if err != nil || got != expected {
		t.Fatalf("CreateUser = %#v, %v", got, err)
	}
	if want := []string{"u", "User", "pw"}; !equalStrings(store.createArgs, want) {
		t.Fatalf("CreateUser args = %#v", store.createArgs)
	}
	got, err = service.UpdateUser(7, "updated", "User")
	if err != nil || got != expected {
		t.Fatalf("UpdateUser = %#v, %v", got, err)
	}
	if store.updateArgs[0] != int64(7) {
		t.Fatalf("UpdateUser id = %#v", store.updateArgs[0])
	}
}

func TestServiceCreateAndUpdateStopOnStoreError(t *testing.T) {
	storeErr := errors.New("store failed")
	store := &mockStore{err: storeErr}
	service := userService.NewService(store)
	if _, err := service.CreateUser("u", "n", "p"); !errors.Is(err, storeErr) {
		t.Fatalf("CreateUser error = %v", err)
	}
	if len(store.findCalls) != 0 {
		t.Fatal("FindUserByID should not be called after create failure")
	}
	if _, err := service.UpdateUser(1, "u", "n"); !errors.Is(err, storeErr) {
		t.Fatalf("UpdateUser error = %v", err)
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
