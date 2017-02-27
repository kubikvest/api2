package app

import "database/sql"

type DB struct {
	*sql.DB
}

type Tx struct {
	*sql.Tx
}

func Open(dataSourceName string) (*DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Begin() (*Tx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{tx}, nil
}
