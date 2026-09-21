// Package race (infrastructure)
package race

import (
	"database/sql"
	"time"

	racingModel "tmp-app-backend/model/racing"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRace(scanner interface{ Scan(...any) error }) (racingModel.Race, error) {
	var race racingModel.Race
	var raceCourse racingModel.Racecourse
	var startTime string

	err := scanner.Scan(
		&race.ID,
		&race.RaceDate,
		&raceCourse.ID,
		&raceCourse.Code,
		&raceCourse.Name,
		&race.RaceNumber,
		&race.RaceName,
		&startTime,
		&race.Surface,
		&race.Distance,
		&race.Direction,
		&race.Weather,
		&race.TrackCondition,
		&race.RaceConditions,
	)
	if err != nil {
		return race, err
	}

	race.StartTime, err = time.Parse("15:04:05", startTime)
	if err != nil {
		return race, err
	}
	race.Racecourse = &raceCourse

	return race, err
}

func (s *Store) FetchRaces() ([]racingModel.Race, error) {
	rows, err := s.db.Query(`
		SELECT
			r.id,
			r.race_date,
			rc.id,
			rc.code,
			rc.name,
			r.race_number,
			r.race_name,
			r.start_time,
			r.surface,
			r.distance,
			r.direction,
			r.weather,
			r.track_condition,
			r.race_conditions
		FROM races r
		INNER JOIN race_courses rc ON rc.id = r.race_course_id
		ORDER BY r.race_date, r.race_number, r.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var races []racingModel.Race

	for rows.Next() {
		race, err := scanRace(rows)
		if err != nil {
			return nil, err
		}

		races = append(races, race)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return races, nil
}

func (s *Store) FindRaceByID(id int64) (racingModel.Race, error) {
	row := s.db.QueryRow(`
		SELECT
			r.id,
			r.race_date,
			rc.id,
			rc.code,
			rc.name,
			r.race_number,
			r.race_name,
			r.start_time,
			r.surface,
			r.distance,
			r.direction,
			r.weather,
			r.track_condition,
			r.race_conditions
		FROM races r
		INNER JOIN race_courses rc ON rc.id = r.race_course_id
		WHERE r.id = ?
	`, id)

	return scanRace(row)
}

func (s *Store) CreateRace(input racingModel.RaceInput) (int64, error) {
	result, err := s.db.Exec(`
		INSERT INTO races (
			race_date,
			race_course_id,
			race_number,
			race_name,
			start_time,
			surface,
			distance,
			direction,
			weather,
			track_condition,
			race_conditions
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		input.RaceDate,
		input.RaceCourseID,
		input.RaceNumber,
		input.RaceName,
		input.StartTime,
		input.Surface,
		input.Distance,
		input.Direction,
		input.Weather,
		input.TrackCondition,
		input.RaceConditions,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (s *Store) UpdateRace(id int64, input racingModel.RaceInput) error {
	result, err := s.db.Exec(`
		UPDATE races
		SET
			race_date = ?,
			race_course_id = ?,
			race_number = ?,
			race_name = ?,
			start_time = ?,
			surface = ?,
			distance = ?,
			direction = ?,
			weather = ?,
			track_condition = ?,
			race_conditions = ?
		WHERE id = ?
	`,
		input.RaceDate,
		input.RaceCourseID,
		input.RaceNumber,
		input.RaceName,
		input.StartTime,
		input.Surface,
		input.Distance,
		input.Direction,
		input.Weather,
		input.TrackCondition,
		input.RaceConditions,
		id,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		_, err := s.FindRaceByID(id)
		return err
	}

	return nil
}

func (s *Store) DeleteRace(id int64) error {
	result, err := s.db.Exec(`DELETE FROM races WHERE id = ?`, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		_, err := s.FindRaceByID(id)
		return err
	}

	return nil
}
