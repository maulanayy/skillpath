package app

import (
	"skillpath/model"
)

type Usecase interface {
	Login(req model.Login) (model.UserInfo, error)
	Create(req model.User) error
}
