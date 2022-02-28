package model

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID        int64       `db:"id" json:"id"`
	Email     string      `db:"email" json:"email"`
	Password  string      `db:"password" json:"password"`
	Token     string      `db:"token" json:"token"`
	Info      Info        `db:"info" json:"info"`
	CreatedAt time.Time   `db:"created_at" json:"created_at"`
	CreatedBy string      `db:"created_by" json:"created_by"`
	UpdatedAt pq.NullTime `db:"updated_at" json:"-"`
	DeletedAt pq.NullTime `db:"deleted_at" json:"-"`
}

type Info struct {
	Name               string `json:"name"`
	TokenGeneratedTime int64  `json:"token_generated_time,omitempty"`
}

type UserInfo struct {
	ID    int64  `db:"id" json:"id"`
	Token string `db:"token" json:"token"`
	Email string `db:"email" json:"email"`
	Name  string `db:"name" json:"name"`
}

type Register struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Name     string `json:"name"`
}
