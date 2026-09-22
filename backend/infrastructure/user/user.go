// Package user (infrastructure)
package user

import (
	"database/sql"

	userModel "keiba-app-backend/model/user"
)

type Store struct {
	db *sql.DB
}

// NewStore データベースを使うユーザーStoreを生成する
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// FetchUsers ユーザー一覧を取得するメソッド
func (s *Store) FetchUsers() ([]userModel.User, error) {
	rows, err := s.db.Query(`
		SELECT id, code, name, password
		FROM users
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []userModel.User

	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// FindUserByCodeAndPass CodeとPasswordを指定してユーザーを取得するメソッド
func (s *Store) FindUserByCodeAndPass(
	code string,
	password string,
) (userModel.User, error) {
	row := s.db.QueryRow(`
		SELECT id, code, name, password
		FROM users
		WHERE code = ? AND password = ?
	`, code, password)

	return scanUser(row)
}

// FindUserByID IDを指定してユーザーを取得するメソッド
func (s *Store) FindUserByID(id int64) (userModel.User, error) {
	row := s.db.QueryRow(`
		SELECT id, code, name, password
		FROM users
		WHERE id = ?
	`, id)
	return scanUser(row)
}

// CreateUser ユーザーを作成するメソッド
func (s *Store) CreateUser(code, name, password string) (int64, error) {
	result, err := s.db.Exec(`
		INSERT INTO users (code, name, password)
		VALUES (?, ?, ?)
	`, code, name, password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// UpdateUser ユーザーを更新するメソッド
func (s *Store) UpdateUser(id int64, code, name string) error {
	result, err := s.db.Exec(`
		UPDATE users
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
		_, err := s.FindUserByID(id)
		return err
	}
	return nil
}

// DeleteUser ユーザーを削除するメソッド
func (s *Store) DeleteUser(id int64) error {
	result, err := s.db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		_, err := s.FindUserByID(id)
		return err
	}
	return nil
}

// scanUser SQLの結果からユーザーモデルを生成する
func scanUser(scanner interface{ Scan(...any) error }) (userModel.User, error) {
	var user userModel.User

	err := scanner.Scan(
		&user.ID,
		&user.Code,
		&user.Name,
		&user.Password,
	)

	return user, err
}
