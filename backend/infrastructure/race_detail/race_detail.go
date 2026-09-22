// Package racedetail (infrastructure)
package racedetail

import (
	"database/sql"
	"time"

	racingModel "keiba-app-backend/model/racing"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRaceDetail(scanner interface{ Scan(...any) error }) (racingModel.RaceDetail, error) {
	var detail racingModel.RaceDetail
	var race racingModel.Race
	var raceCourse racingModel.Racecourse
	var startTime string

	err := scanner.Scan(
		&detail.ID,
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
		&detail.HorseNumber,
		&detail.FrameNumber,
		&detail.HorseName,
		&detail.Sex,
		&detail.Age,
		&detail.Weight,
		&detail.Jockey,
		&detail.Stable,
		&detail.BodyWeight,
		&detail.BodyWeightChange,
		&detail.Odds,
		&detail.Popularity,
	)
	if err != nil {
		return detail, err
	}

	race.StartTime, err = time.Parse("15:04:05", startTime)
	if err != nil {
		return detail, err
	}
	race.Racecourse = &raceCourse
	detail.Race = &race

	return detail, nil
}

const raceDetailSelect = `
	SELECT
		rd.id,
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
		r.race_conditions,
		rd.horse_number,
		rd.frame_number,
		rd.horse_name,
		rd.sex,
		rd.age,
		rd.weight,
		rd.jockey,
		rd.stable,
		rd.body_weight,
		rd.body_weight_change,
		rd.odds,
		rd.popularity
	FROM race_details rd
	INNER JOIN races r ON r.id = rd.race_id
	INNER JOIN race_courses rc ON rc.id = r.race_course_id
`

func (s *Store) FetchRaceDetails(raceID int64) ([]racingModel.RaceDetail, error) {
	rows, err := s.db.Query(raceDetailSelect+`
		WHERE rd.race_id = ?
		ORDER BY rd.horse_number
	`, raceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []racingModel.RaceDetail
	for rows.Next() {
		detail, err := scanRaceDetail(rows)
		if err != nil {
			return nil, err
		}
		details = append(details, detail)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return details, nil
}

func (s *Store) FindRaceDetailByID(raceID, detailID int64) (racingModel.RaceDetail, error) {
	row := s.db.QueryRow(raceDetailSelect+`
		WHERE rd.race_id = ? AND rd.id = ?
	`, raceID, detailID)

	return scanRaceDetail(row)
}
