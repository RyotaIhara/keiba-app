// Package user (Service)
package user

import userModel "keiba-app-backend/model/user"

type Service struct {
	store userStore
}

type userStore interface {
	FetchUsers() ([]userModel.User, error)
	FindUserByID(id int64) (userModel.User, error)
	CreateUser(code, name, password string) (int64, error)
	UpdateUser(id int64, code, name string) error
	DeleteUser(id int64) error
}

// NewService ユーザーServiceを生成する
func NewService(store userStore) *Service {
	return &Service{store: store}
}

// GetUsers ユーザー一覧を取得する
func (s *Service) GetUsers() ([]userModel.User, error) {
	return s.store.FetchUsers()
}

// GetUser 指定されたIDのユーザーを取得する
func (s *Service) GetUser(id int64) (userModel.User, error) {
	return s.store.FindUserByID(id)
}

// CreateUser ユーザーを作成する
func (s *Service) CreateUser(code, name, password string) (userModel.User, error) {
	id, err := s.store.CreateUser(code, name, password)
	if err != nil {
		return userModel.User{}, err
	}
	return s.store.FindUserByID(id)
}

// UpdateUser 指定されたIDのユーザーを更新する
func (s *Service) UpdateUser(id int64, code, name string) (userModel.User, error) {
	if err := s.store.UpdateUser(id, code, name); err != nil {
		return userModel.User{}, err
	}
	return s.store.FindUserByID(id)
}

// DeleteUser 指定されたIDのユーザーを削除する
func (s *Service) DeleteUser(id int64) error {
	return s.store.DeleteUser(id)
}
