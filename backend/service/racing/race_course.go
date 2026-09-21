package racing

import (
	raceCourseInfrastructure "tmp-app-backend/infrastructure/race_course"
	racingModel "tmp-app-backend/model/racing"
)

type RaceCourseService struct {
	repository *raceCourseInfrastructure.Repository
}

func NewRaceCourseService(repository *raceCourseInfrastructure.Repository) *RaceCourseService {
	return &RaceCourseService{repository: repository}
}

func (s *RaceCourseService) GetRaceCourses() ([]racingModel.Racecourse, error) {
	return s.repository.FetchRaceCourses()
}
