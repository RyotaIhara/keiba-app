// Package user (Service)
package user

import (
	userInfrastructure "tmp-app-backend/infrastructure/user"
	userModel "tmp-app-backend/model/user"
)

type Service struct {
	repository *userInfrastructure.Repository
}

func NewService(repository *userInfrastructure.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetUsers() ([]userModel.User, error) {
	return s.repository.FetchUsers()
}
