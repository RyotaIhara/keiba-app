// Package race (infrastructure)
package race

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

func scanRace(scanner interface{ Scan(...any) error }) (racingModel.Race, error) {
	var race racingModel.Race
	var raceCourse racingModel.Racecourse

	err := scanner.Scan(
		&race.ID,
		&race.RaceDate,
		&raceCourse.ID,
		&raceCourse.Code,
		&raceCourse.Name,
		&race.RaceNumber,
		&race.RaceName,
		&race.StartTime,
		&race.Surface,
		&race.Distance,
		&race.Direction,
		&race.Weather,
		&race.TrackCondition,
		&race.RaceConditions,
	)
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
