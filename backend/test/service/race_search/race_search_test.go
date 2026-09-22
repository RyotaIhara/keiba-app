package racesearch

import (
	"errors"
	"testing"

	raceModel "keiba-app-backend/model/race"
	raceSupport "keiba-app-backend/model/race/support"
	raceSearchService "keiba-app-backend/service/race_search"
)

type mockStore struct {
	races []raceModel.Race
	err   error
}

func (m *mockStore) SearchRaces(raceSupport.RaceSearchInput) ([]raceModel.Race, error) {
	return m.races, m.err
}

func TestSearchServiceDelegatesSearch(t *testing.T) {
	expected := []raceModel.Race{{ID: 1}}
	got, err := raceSearchService.NewSearchService(&mockStore{races: expected}).SearchRaces(raceSupport.RaceSearchInput{})
	if err != nil || len(got) != 1 || got[0] != expected[0] {
		t.Fatalf("SearchRaces = %#v, %v", got, err)
	}
}

func TestSearchServicePropagatesError(t *testing.T) {
	errExpected := errors.New("store failed")
	_, err := raceSearchService.NewSearchService(&mockStore{err: errExpected}).SearchRaces(raceSupport.RaceSearchInput{})
	if !errors.Is(err, errExpected) {
		t.Fatalf("error = %v", err)
	}
}
