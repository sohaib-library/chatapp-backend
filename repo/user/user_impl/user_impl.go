package user_impl

import (
	"github.com/sohaib-library/chatapp-backend/repo/user"

	"gorm.io/gorm"
)

var _ user.UserRepo = (*UserImpl)(nil)

type UserImpl struct {
	db *gorm.DB
}

func NewUserImpl(db *gorm.DB) user.UserRepo {
	return &UserImpl{db: db}
}
