// Package race (Service)
package race

import (
	raceInfrastructure "keiba-app-backend/infrastructure/race"
	raceModel "keiba-app-backend/model/race"
	raceSupport "keiba-app-backend/model/race/support"
)

type RaceService struct {
	store *raceInfrastructure.Store
}

func NewRaceService(store *raceInfrastructure.Store) *RaceService {
	return &RaceService{store: store}
}

func (s *RaceService) GetRaces() ([]raceModel.Race, error) {
	return s.store.FetchRaces()
}

func (s *RaceService) GetRace(id int64) (raceModel.Race, error) {
	return s.store.FindRaceByID(id)
}

func (s *RaceService) GetRaceDetails(raceID int64) ([]raceModel.RaceDetail, error) {
	return s.store.FetchRaceDetailsByRaceID(raceID)
}

func (s *RaceService) CreateRace(input raceSupport.RaceInput) (raceModel.Race, error) {
	id, err := s.store.CreateRace(input)
	if err != nil {
		return raceModel.Race{}, err
	}

	return s.store.FindRaceByID(id)
}

func (s *RaceService) UpdateRace(id int64, input raceSupport.RaceInput) (raceModel.Race, error) {
	if err := s.store.UpdateRace(id, input); err != nil {
		return raceModel.Race{}, err
	}

	return s.store.FindRaceByID(id)
}

func (s *RaceService) DeleteRace(id int64) error {
	return s.store.DeleteRace(id)
}
