// Package database provides the application's database connection.
package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Connect 環境変数を使ってデータベースへ接続する
func Connect() (*sql.DB, error) {
	host, err := requiredEnv("DB_HOST")
	if err != nil {
		return nil, err
	}
	port, err := requiredEnv("DB_PORT")
	if err != nil {
		return nil, err
	}
	name, err := requiredEnv("DB_NAME")
	if err != nil {
		return nil, err
	}
	user, err := requiredEnv("DB_USER")
	if err != nil {
		return nil, err
	}
	password, err := requiredEnv("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	config := mysql.Config{
		User:      user,
		Passwd:    password,
		Net:       "tcp",
		Addr:      host + ":" + port,
		DBName:    name,
		ParseTime: true,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}

// requiredEnv 指定された必須環境変数の値を取得する
func requiredEnv(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", fmt.Errorf("%s is required", name)
	}

	return value, nil
}
