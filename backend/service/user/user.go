// Package user (Service)
package user

import (
	userInfrastructure "tmp-app-backend/infrastructure/user"
	userModel "tmp-app-backend/model/user"
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
