package auth_impl

import (
	"github.com/sohaib-library/chatapp-backend/repo/auth"
	authService "github.com/sohaib-library/chatapp-backend/service/auth"
)

var _ authService.AuthService = (*AuthServiceImpl)(nil)

type AuthServiceImpl struct {
	repo auth.AuthRepo
}

func NewAuth(repo auth.AuthRepo) authService.AuthService {
	return &AuthServiceImpl{repo: repo}
}
