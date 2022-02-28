package db

import (
	"skillpath/model"
)

type User interface {
	Create(body *model.User) error
	GetByEmail(email string) (model.User, error)
}
