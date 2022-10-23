package postgres

import "github.com/jmoiron/sqlx"

func NewMock() (conn *sqlx.DB, err error) {
	config := Config{
		Host:     "localhost",
		Port:     "5433",
		Username: "postgres",
		Password: "pass",
		Database: "chat",
	}

	return New(config)
}
