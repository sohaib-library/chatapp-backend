package auth_impl

import (
	authHandler "github.com/sohaib-library/chatapp-backend/handler/auth"
	authService "github.com/sohaib-library/chatapp-backend/service/auth"
)

var _ authHandler.AuthHandler = (*Handler)(nil)

type Handler struct {
	Authuser authService.AuthService
}

func NewHandler(authSvc authService.AuthService) *Handler {
	return &Handler{Authuser: authSvc}
}
