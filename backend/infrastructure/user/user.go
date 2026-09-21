// Package user (infrastructure)
package user

import (
	"database/sql"

	userModel "tmp-app-backend/model/user"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

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

func (r *Repository) FetchUsers() ([]userModel.User, error) {
	rows, err := r.db.Query(`
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

func (r *Repository) FindUserByCodeAndPass(
	code string,
	password string,
) (userModel.User, error) {
	row := r.db.QueryRow(`
		SELECT id, code, name, password
		FROM users
		WHERE code = ? AND password = ?
	`, code, password)

	return scanUser(row)
}
