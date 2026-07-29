package auth_impl

import (
	"chatapp-backend/repo/auth"

	"gorm.io/gorm"
)

var _ auth.AuthRepo = (*AuthImpl)(nil)

type AuthImpl struct {
	db *gorm.DB
}

func NewAuthImpl(db *gorm.DB) auth.AuthRepo {
	return &AuthImpl{db: db}
}
