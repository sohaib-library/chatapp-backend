package user_impl

import (
	"chatapp-backend/repo/user"
	userService "chatapp-backend/service/user"
)

var _ userService.UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	repo user.UserRepo
}

func NewUser(repo user.UserRepo) userService.UserService {
	return &UserServiceImpl{repo: repo}
}
