// Package user (Service)
package user

import (
	userInfrastructure "keiba-app-backend/infrastructure/user"
	userModel "keiba-app-backend/model/user"
)

type Service struct {
	store *userInfrastructure.Store
}

func NewService(store *userInfrastructure.Store) *Service {
	return &Service{store: store}
}

func (s *Service) GetUsers() ([]userModel.User, error) {
	return s.store.FetchUsers()
}

func (s *Service) GetUser(id int64) (userModel.User, error) {
	return s.store.FindUserByID(id)
}

func (s *Service) CreateUser(code, name, password string) (userModel.User, error) {
	id, err := s.store.CreateUser(code, name, password)
	if err != nil {
		return userModel.User{}, err
	}
	return s.store.FindUserByID(id)
}

func (s *Service) UpdateUser(id int64, code, name string) (userModel.User, error) {
	if err := s.store.UpdateUser(id, code, name); err != nil {
		return userModel.User{}, err
	}
	return s.store.FindUserByID(id)
}

func (s *Service) DeleteUser(id int64) error {
	return s.store.DeleteUser(id)
}
