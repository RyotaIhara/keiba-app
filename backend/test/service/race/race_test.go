package race

import (
	"errors"
	"testing"

	raceModel "keiba-app-backend/model/race"
	raceSupport "keiba-app-backend/model/race/support"
	raceService "keiba-app-backend/service/race"
)

type mockStore struct {
	races   []raceModel.Race
	details []raceModel.RaceDetail
	race    raceModel.Race
	id      int64
	err     error
}

func (m *mockStore) FetchRaces() ([]raceModel.Race, error)      { return m.races, m.err }
func (m *mockStore) FindRaceByID(int64) (raceModel.Race, error) { return m.race, m.err }
func (m *mockStore) FetchRaceDetailsByRaceID(int64) ([]raceModel.RaceDetail, error) {
	return m.details, m.err
}
func (m *mockStore) CreateRace(raceSupport.RaceInput) (int64, error) { return m.id, m.err }
func (m *mockStore) UpdateRace(int64, raceSupport.RaceInput) error   { return m.err }
func (m *mockStore) DeleteRace(int64) error                          { return m.err }

func TestRaceServiceDelegatesAllMethods(t *testing.T) {
	expected := raceModel.Race{ID: 3, RaceName: "記念"}
	store := &mockStore{races: []raceModel.Race{expected}, race: expected, details: []raceModel.RaceDetail{{ID: 4}}, id: 3}
	service := raceService.NewRaceService(store)
	if got, err := service.GetRaces(); err != nil || got[0] != expected {
		t.Fatalf("GetRaces = %#v, %v", got, err)
	}
	if got, err := service.GetRace(3); err != nil || got != expected {
		t.Fatalf("GetRace = %#v, %v", got, err)
	}
	if got, err := service.GetRaceDetails(3); err != nil || got[0].ID != 4 {
		t.Fatalf("GetRaceDetails = %#v, %v", got, err)
	}
	if got, err := service.CreateRace(raceSupport.RaceInput{}); err != nil || got != expected {
		t.Fatalf("CreateRace = %#v, %v", got, err)
	}
	if got, err := service.UpdateRace(3, raceSupport.RaceInput{}); err != nil || got != expected {
		t.Fatalf("UpdateRace = %#v, %v", got, err)
	}
	if err := service.DeleteRace(3); err != nil {
		t.Fatalf("DeleteRace = %v", err)
	}
}

func TestRaceServicePropagatesErrors(t *testing.T) {
	storeErr := errors.New("store failed")
	service := raceService.NewRaceService(&mockStore{err: storeErr})
	if _, err := service.GetRaces(); !errors.Is(err, storeErr) {
		t.Fatal(err)
	}
	if _, err := service.CreateRace(raceSupport.RaceInput{}); !errors.Is(err, storeErr) {
		t.Fatal(err)
	}
	if _, err := service.UpdateRace(1, raceSupport.RaceInput{}); !errors.Is(err, storeErr) {
		t.Fatal(err)
	}
}
