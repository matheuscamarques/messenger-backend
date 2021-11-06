package main

import (
	"github.com/go-pg/pg/v10"
)

var db *pg.DB = pg.Connect(&pg.Options{
	User:     "postgres",
	Password: "postgres",
	Database: "postgres",
	Addr:     "localhost:5432",
})

func DBConnection() *pg.DB {
	return db
}
