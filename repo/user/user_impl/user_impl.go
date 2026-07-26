package user_impl

import (
	"database/sql"

	"chatapp-backend/repo/user"
)

type UserImpl struct {
	db *sql.DB
}

func NewUserImpl(db *sql.DB) user.UserRepo {
	return &UserImpl{db: db}
}
