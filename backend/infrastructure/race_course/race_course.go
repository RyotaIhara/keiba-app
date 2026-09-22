// Package racecourse (infrastructure)
package racecourse

import (
	"database/sql"

	raceCourseModel "keiba-app-backend/model/race_course"
)

type Store struct {
	db *sql.DB
}

// NewStore データベースを使う競馬場Storeを生成する
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// FetchRaceCourses 競馬場一覧取得するメソッド
func (s *Store) FetchRaceCourses() ([]raceCourseModel.Racecourse, error) {
	rows, err := s.db.Query(`
		SELECT id, code, name
		FROM race_courses
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var raceCourses []raceCourseModel.Racecourse

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

// FindRaceCourseByID IDを指定して競馬場を取得するメソッド
func (s *Store) FindRaceCourseByID(id int64) (raceCourseModel.Racecourse, error) {
	row := s.db.QueryRow(`
		SELECT id, code, name
		FROM race_courses
		WHERE id = ?
	`, id)
	return scanRaceCourse(row)
}

// CreateRaceCourse 競馬場を作成するメソッド
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

// UpdateRaceCourse 競馬場を更新するメソッド
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

// DeleteRaceCourse 競馬場を削除するメソッド
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

// scanRaceCourse SQLの結果から競馬場モデルを生成する
func scanRaceCourse(scanner interface{ Scan(...any) error }) (raceCourseModel.Racecourse, error) {
	var raceCourse raceCourseModel.Racecourse

	err := scanner.Scan(
		&raceCourse.ID,
		&raceCourse.Code,
		&raceCourse.Name,
	)

	return raceCourse, err
}
