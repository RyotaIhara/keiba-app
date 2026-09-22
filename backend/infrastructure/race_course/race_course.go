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

func (s *Store) FindRaceCourseByID(id int64) (racingModel.Racecourse, error) {
	row := s.db.QueryRow(`
		SELECT id, code, name
		FROM race_courses
		WHERE id = ?
	`, id)
	return scanRaceCourse(row)
}

func (s *Store) CreateRaceCourse(code, name string) (int64, error) {
	result, err := s.db.Exec(`
		INSERT INTO race_courses (code, name)
		VALUES (?, ?)
	`, code, name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (s *Store) UpdateRaceCourse(id int64, code, name string) error {
	result, err := s.db.Exec(`
		UPDATE race_courses
		SET code = ?, name = ?
		WHERE id = ?
	`, code, name, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		_, err := s.FindRaceCourseByID(id)
		return err
	}
	return nil
}

func (s *Store) DeleteRaceCourse(id int64) error {
	result, err := s.db.Exec(`DELETE FROM race_courses WHERE id = ?`, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		_, err := s.FindRaceCourseByID(id)
		return err
	}
	return nil
}
