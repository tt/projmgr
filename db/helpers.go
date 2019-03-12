package db

import (
	"github.com/jmoiron/sqlx"
)

func insert(tx *sqlx.Tx, dest interface{}, query string, arg interface{}) error {
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return err
	}

	return stmt.Get(dest, arg)
}

func get(db *sqlx.DB, dest interface{}, query string, args ...interface{}) error {
	return mapError(db.Get(dest, query, args...))
}

func list(db *sqlx.DB, dest interface{}, query string, args ...interface{}) error {
	return db.Select(dest, query, args...)
}
