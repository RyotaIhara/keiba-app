// Package race (Service)
package race

import (
	raceModel "keiba-app-backend/model/race"
	raceSupport "keiba-app-backend/model/race/support"
)

type RaceService struct {
	store raceStore
}

type raceStore interface {
	FetchRaces() ([]raceModel.Race, error)
	FindRaceByID(id int64) (raceModel.Race, error)
	FetchRaceDetailsByRaceID(raceID int64) ([]raceModel.RaceDetail, error)
	CreateRace(input raceSupport.RaceInput) (int64, error)
	UpdateRace(id int64, input raceSupport.RaceInput) error
	DeleteRace(id int64) error
}

// NewRaceService レースServiceを生成する
func NewRaceService(store raceStore) *RaceService {
	return &RaceService{store: store}
}

// GetRaces レース一覧を取得する
func (s *RaceService) GetRaces() ([]raceModel.Race, error) {
	return s.store.FetchRaces()
}

// GetRace 指定されたIDのレースを取得する
func (s *RaceService) GetRace(id int64) (raceModel.Race, error) {
	return s.store.FindRaceByID(id)
}

// GetRaceDetails 指定されたレースの出走馬一覧を取得する
func (s *RaceService) GetRaceDetails(raceID int64) ([]raceModel.RaceDetail, error) {
	return s.store.FetchRaceDetailsByRaceID(raceID)
}

// CreateRace レースを作成する
func (s *RaceService) CreateRace(input raceSupport.RaceInput) (raceModel.Race, error) {
	id, err := s.store.CreateRace(input)
	if err != nil {
		return raceModel.Race{}, err
	}

	return s.store.FindRaceByID(id)
}

// UpdateRace 指定されたIDのレースを更新する
func (s *RaceService) UpdateRace(id int64, input raceSupport.RaceInput) (raceModel.Race, error) {
	if err := s.store.UpdateRace(id, input); err != nil {
		return raceModel.Race{}, err
	}

	return s.store.FindRaceByID(id)
}

// DeleteRace 指定されたIDのレースを削除する
func (s *RaceService) DeleteRace(id int64) error {
	return s.store.DeleteRace(id)
}
