// Package racing (Service)
package racing

import (
	raceInfrastructure "tmp-app-backend/infrastructure/race"
	racingModel "tmp-app-backend/model/racing"
)

type RaceService struct {
	store *raceInfrastructure.Store
}

func NewRaceService(store *raceInfrastructure.Store) *RaceService {
	return &RaceService{store: store}
}

func (s *RaceService) GetRaces() ([]racingModel.Race, error) {
	return s.store.FetchRaces()
}
