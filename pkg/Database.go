package pkg

import (
	"context"
	"fmt"
	"skillpath/model"

	pg "github.com/go-pg/pg/v10"
)

func NewDatabase(ctx context.Context, params model.DBParams) (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", params.Address, params.Port),
		Database: params.DBName,
		User:     params.User,
		Password: params.Password,
	})

	return db, nil
}
