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

func (s *RaceService) GetRace(id int64) (racingModel.Race, error) {
	return s.store.FindRaceByID(id)
}

func (s *RaceService) CreateRace(input racingModel.RaceInput) (racingModel.Race, error) {
	id, err := s.store.CreateRace(input)
	if err != nil {
		return racingModel.Race{}, err
	}

	return s.store.FindRaceByID(id)
}

func (s *RaceService) UpdateRace(id int64, input racingModel.RaceInput) (racingModel.Race, error) {
	if err := s.store.UpdateRace(id, input); err != nil {
		return racingModel.Race{}, err
	}

	return s.store.FindRaceByID(id)
}

func (s *RaceService) DeleteRace(id int64) error {
	return s.store.DeleteRace(id)
}
