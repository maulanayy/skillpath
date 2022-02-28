package app

import (
	"skillpath/model"
)

type Middleware struct {
	Config  *model.Config
	Usecase Usecase
}

// New to create new middleware
func NewMiddleware(uc Usecase, cfg *model.Config) *Middleware {
	return &Middleware{
		Config:  cfg,
		Usecase: uc,
	}
}
