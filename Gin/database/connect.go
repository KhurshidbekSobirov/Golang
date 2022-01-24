package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Conn() *sql.DB{
	db,err := sql.Open("postgres","dbname=firstdb user=khurshid password=X sslmode=disable")
	if err != nil{
		panic(err)
	}
	return db
}