package auth_impl

import (
	"database/sql"

	"chatapp-backend/repo/auth"
)

type AuthImpl struct {
	db *sql.DB
}

func NewAuthImpl(db *sql.DB) auth.AuthRepo {
	return &AuthImpl{db: db}
}
