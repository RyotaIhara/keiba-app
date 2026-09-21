package racing

import (
	raceCourseInfrastructure "tmp-app-backend/infrastructure/race_course"
	racingModel "tmp-app-backend/model/racing"
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
