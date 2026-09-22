// Package race (infrastructure)
package race

import (
	"database/sql"
	"strings"
	"time"

	raceModel "keiba-app-backend/model/race"
	raceSupport "keiba-app-backend/model/race/support"
	raceCourseModel "keiba-app-backend/model/race_course"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRace(scanner interface{ Scan(...any) error }) (raceModel.Race, error) {
	var race raceModel.Race
	var raceCourse raceCourseModel.Racecourse
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

func scanRaceDetail(scanner interface{ Scan(...any) error }) (raceModel.RaceDetail, error) {
	var detail raceModel.RaceDetail
	var race raceModel.Race
	var raceCourse raceCourseModel.Racecourse
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

// FetchRaces レース一覧取得するメソッド
func (s *Store) FetchRaces() ([]raceModel.Race, error) {
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

	var races []raceModel.Race

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

// SearchRaces 検索条件を指定してレース一覧を取得するメソッド
func (s *Store) SearchRaces(input raceSupport.RaceSearchInput) ([]raceModel.Race, error) {
	query := `
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
	`
	conditions := make([]string, 0, 2)
	args := make([]any, 0, 2)

	if input.RaceDate != nil {
		conditions = append(conditions, "r.race_date = ?")
		args = append(args, *input.RaceDate)
	}
	if input.RaceCourseID != nil {
		conditions = append(conditions, "r.race_course_id = ?")
		args = append(args, *input.RaceCourseID)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += " ORDER BY r.race_date, r.race_number, r.id"

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var races []raceModel.Race
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

// FindRaceByID IDを指定してレース情報を取得するメソッド
func (s *Store) FindRaceByID(id int64) (raceModel.Race, error) {
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

// FetchRaceDetailsByRaceID IDを指定してDetails一覧を取得するメソッド
func (s *Store) FetchRaceDetailsByRaceID(raceID int64) ([]raceModel.RaceDetail, error) {
	rows, err := s.db.Query(`
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
		WHERE rd.race_id = ?
		ORDER BY rd.horse_number
	`, raceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []raceModel.RaceDetail
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

// CreateRace レース情報を作成する
func (s *Store) CreateRace(input raceSupport.RaceInput) (int64, error) {
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

// UpdateRace レース情報を更新する
func (s *Store) UpdateRace(id int64, input raceSupport.RaceInput) error {
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

// DeleteRace レース情報を削除する
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
