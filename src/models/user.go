package models

import (
	"github.com/matheuscamarques/messenger-backend/src/main"
)

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"username"`
}

func (u *User) GetByID(id int64) (*User, error) {
	user := &User{}
	err := main.DBConnection().Model(user).Where("id = ?", id).Select()
	return user, err
}

func (u *User) GetByUsername(username string) (*User, error) {
	err := main.DBConnection().Model(u).Where("username = ?", username).Select()
	return u, err
}

func (u *User) Create() error {
	_, err := main.DBConnection().Model(u).Insert()
	return err
}
