package database

import (
	_"github.com/lib/pq"
	"database/sql"
)

func Con() *sql.DB {
	db, err := sql.Open("postgres", "dbname=firstdb  user=khurshid password=X sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
