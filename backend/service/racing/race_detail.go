// Package racing (Service)
package racing

import (
	raceDetailInfrastructure "keiba-app-backend/infrastructure/race_detail"
	racingModel "keiba-app-backend/model/racing"
)

type RaceDetailService struct {
	store *raceDetailInfrastructure.Store
}

func NewRaceDetailService(store *raceDetailInfrastructure.Store) *RaceDetailService {
	return &RaceDetailService{store: store}
}

func (s *RaceDetailService) GetRaceDetails(raceID int64) ([]racingModel.RaceDetail, error) {
	return s.store.FetchRaceDetails(raceID)
}

func (s *RaceDetailService) GetRaceDetail(raceID, detailID int64) (racingModel.RaceDetail, error) {
	return s.store.FindRaceDetailByID(raceID, detailID)
}
