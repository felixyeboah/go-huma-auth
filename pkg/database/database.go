package database

import (
	"database/sql"
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

var Database *sql.DB

func Connect(db string) (*sql.DB, error) {
	dbase, err := newDB(db)
	if err != nil {
		return nil, err
	}

	Database = dbase

	return dbase, nil
}
