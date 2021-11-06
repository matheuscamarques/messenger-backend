package models

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"username"`
}

func (u *User) GetByID(id int64) (*User, error) {
	user := &User{}
	err := db.Model(user).Where("id = ?", id).Select()
	return user, err
}

func (u *User) GetByUsername(username string) (*User, error) {
	err := db.Model(u).Where("username = ?", username).Select()
	return u, err
}

func (u *User) Create() error {
	_, err := db.Model(u).Insert()
	return err
}
