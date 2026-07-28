package auth_impl

import (
	"database/sql"

	"chatapp-backend/repo/auth"
)

var _ auth.AuthRepo = (*AuthImpl)(nil)

type AuthImpl struct {
	db *sql.DB
}

func NewAuthImpl(db *sql.DB) auth.AuthRepo {
	return &AuthImpl{db: db}
}
