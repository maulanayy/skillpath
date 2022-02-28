package app

import (
	"skillpath/db"
	"skillpath/repository"

	"github.com/go-pg/pg/v10"
)

type Repo struct {
	UserDB db.User
}

func NewRepo(db *pg.DB) *Repo {
	return &Repo{
		UserDB: repository.NewUser(db),
	}
}
