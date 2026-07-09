package auth_impl

import (
	"chatapp-backend/repo/auth"
	authService "chatapp-backend/service/auth"
)

type AuthServiceImpl struct {
	repo auth.AuthRepo
}

func NewAuth(repo auth.AuthRepo) authService.AuthService {
	return &AuthServiceImpl{repo: repo}
}
