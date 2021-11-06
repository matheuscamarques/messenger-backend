package models

import (
	"github.com/go-pg/pg/v10"
)

var db *pg.DB = pg.Connect(&pg.Options{
	User:     "postgres",
	Password: "postgres",
})

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"username"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	return nil
}

func (u *User) Create() error {
	return nil
}
