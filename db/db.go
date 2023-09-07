package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectDb() *sql.DB {
	host := ""
	user := ""
	pass := ""
	dbname := ""

	conexao := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
