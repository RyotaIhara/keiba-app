// Package racesearch (Service)
package racesearch

import (
	raceModel "keiba-app-backend/model/race"
	raceSupport "keiba-app-backend/model/race/support"
)

type SearchService struct {
	store raceSearchStore
}

type raceSearchStore interface {
	SearchRaces(input raceSupport.RaceSearchInput) ([]raceModel.Race, error)
}

// NewSearchService レース検索Serviceを生成する
func NewSearchService(store raceSearchStore) *SearchService {
	return &SearchService{store: store}
}

// SearchRaces 条件を指定してレース一覧を検索する
func (s *SearchService) SearchRaces(input raceSupport.RaceSearchInput) ([]raceModel.Race, error) {
	return s.store.SearchRaces(input)
}
