// Package racecourse (infrastructure)
package racecourse

import (
	"database/sql"

	racingModel "tmp-app-backend/model/racing"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
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

func (r *Repository) FetchRaceCourses() ([]racingModel.Racecourse, error) {
	rows, err := r.db.Query(`
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
