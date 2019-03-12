package db

import (
	"database/sql"

	"github.com/heroku/projmgr"
)

var errorMap = map[error]error{
	sql.ErrNoRows: projmgr.ErrNotFound,
}

func mapError(err error) error {
	v, ok := errorMap[err]
	if ok {
		return v
	}

	return err
}
