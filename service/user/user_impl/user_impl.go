package user_impl

import (
	"github.com/sohaib-library/chatapp-backend/repo/user"
	userService "github.com/sohaib-library/chatapp-backend/service/user"
)

var _ userService.UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	repo user.UserRepo
}

func NewUser(repo user.UserRepo) userService.UserService {
	return &UserServiceImpl{repo: repo}
}
