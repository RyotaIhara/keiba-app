package racecourse

import (
	raceCourseInfrastructure "keiba-app-backend/infrastructure/race_course"
	raceCourseModel "keiba-app-backend/model/race_course"
)

type RaceCourseService struct {
	store *raceCourseInfrastructure.Store
}

func NewRaceCourseService(store *raceCourseInfrastructure.Store) *RaceCourseService {
	return &RaceCourseService{store: store}
}

func (s *RaceCourseService) GetRaceCourses() ([]raceCourseModel.Racecourse, error) {
	return s.store.FetchRaceCourses()
}

func (s *RaceCourseService) GetRaceCourse(id int64) (raceCourseModel.Racecourse, error) {
	return s.store.FindRaceCourseByID(id)
}

func (s *RaceCourseService) CreateRaceCourse(code, name string) (raceCourseModel.Racecourse, error) {
	id, err := s.store.CreateRaceCourse(code, name)
	if err != nil {
		return raceCourseModel.Racecourse{}, err
	}
	return s.store.FindRaceCourseByID(id)
}

func (s *RaceCourseService) UpdateRaceCourse(id int64, code, name string) (raceCourseModel.Racecourse, error) {
	if err := s.store.UpdateRaceCourse(id, code, name); err != nil {
		return raceCourseModel.Racecourse{}, err
	}
	return s.store.FindRaceCourseByID(id)
}

func (s *RaceCourseService) DeleteRaceCourse(id int64) error {
	return s.store.DeleteRaceCourse(id)
}
