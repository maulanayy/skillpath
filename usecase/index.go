package usecase

import (
	"skillpath/app"
	"skillpath/db"
	"skillpath/model"
)

type usecase struct {
	userDB db.User
	cfg    *model.Config
}

func NewUsecase(repo *app.Repo, cfg *model.Config) app.Usecase {
	return &usecase{
		cfg:    cfg,
		userDB: repo.UserDB,
	}
}
