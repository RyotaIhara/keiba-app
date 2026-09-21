// Package racecourse (infrastructure)
package racecourse

import (
	"database/sql"

	racingModel "tmp-app-backend/model/racing"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRaceCourse(scanner interface{ Scan(...any) error }) (racingModel.Racecourse, error) {
	var raceCourse racingModel.Racecourse

	err := scanner.Scan(
		&raceCourse.ID,
		&raceCourse.Code,
		&raceCourse.Name,
	)

	return raceCourse, err
}

func (s *Store) FetchRaceCourses() ([]racingModel.Racecourse, error) {
	rows, err := s.db.Query(`
		SELECT id, code, name
		FROM race_courses
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var raceCourses []racingModel.Racecourse

	for rows.Next() {
		raceCourse, err := scanRaceCourse(rows)
		if err != nil {
			return nil, err
		}

		raceCourses = append(raceCourses, raceCourse)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return raceCourses, nil
}
