// Package racing (Service)
package racing

import (
	raceInfrastructure "tmp-app-backend/infrastructure/race"
	racingModel "tmp-app-backend/model/racing"
)

type RaceService struct {
	repository *raceInfrastructure.Repository
}

func NewRaceService(repository *raceInfrastructure.Repository) *RaceService {
	return &RaceService{repository: repository}
}

func (s *RaceService) GetRaces() ([]racingModel.Race, error) {
	return s.repository.FetchRaces()
}
