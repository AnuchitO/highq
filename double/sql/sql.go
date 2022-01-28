package sql

import "database/sql"

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func execQuery(db DB, query string, args ...interface{}) (int64, error) {
	res, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return ra, nil
}
