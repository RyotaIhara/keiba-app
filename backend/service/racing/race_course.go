package racing

import (
	raceCourseInfrastructure "keiba-app-backend/infrastructure/race_course"
	racingModel "keiba-app-backend/model/racing"
)

type RaceCourseService struct {
	store *raceCourseInfrastructure.Store
}

func NewRaceCourseService(store *raceCourseInfrastructure.Store) *RaceCourseService {
	return &RaceCourseService{store: store}
}

func (s *RaceCourseService) GetRaceCourses() ([]racingModel.Racecourse, error) {
	return s.store.FetchRaceCourses()
}

func (s *RaceCourseService) GetRaceCourse(id int64) (racingModel.Racecourse, error) {
	return s.store.FindRaceCourseByID(id)
}

func (s *RaceCourseService) CreateRaceCourse(code, name string) (racingModel.Racecourse, error) {
	id, err := s.store.CreateRaceCourse(code, name)
	if err != nil {
		return racingModel.Racecourse{}, err
	}
	return s.store.FindRaceCourseByID(id)
}

func (s *RaceCourseService) UpdateRaceCourse(id int64, code, name string) (racingModel.Racecourse, error) {
	if err := s.store.UpdateRaceCourse(id, code, name); err != nil {
		return racingModel.Racecourse{}, err
	}
	return s.store.FindRaceCourseByID(id)
}

func (s *RaceCourseService) DeleteRaceCourse(id int64) error {
	return s.store.DeleteRaceCourse(id)
}
