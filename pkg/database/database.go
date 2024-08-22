package database

import (
	"database/sql"
	"huma-auth/config"
)

func newDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func Connect() (*sql.DB, error) {
	env, err := config.Env()
	if err != nil {
		return nil, err
	}

	dbase, err := newDB(env.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	return dbase, nil
}
