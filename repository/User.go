package repository

import (
	"skillpath/db"
	"skillpath/model"

	"github.com/go-pg/pg/v10"
)

type user struct {
	db *pg.DB
}

func NewUser(db *pg.DB) db.User {
	return &user{
		db: db,
	}
}

func (u *user) Create(body *model.User) error {

	db := u.db.Conn()
	defer db.Close()

	_, err := db.Model(body).Insert()

	return err
}

func (u *user) GetByEmail(email string) (model.User, error) {

	user := new(model.User)

	db := u.db.Conn()
	defer db.Close()

	err := db.Model(user).Where("email = ?", email).Select()

	if err != nil {
		return *user, err
	}

	return *user, nil
}
