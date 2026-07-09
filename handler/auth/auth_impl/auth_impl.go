package auth_impl

import authService "chatapp-backend/service/auth"

type Handler struct {
	Authuser authService.AuthService
}

func NewHandler(authSvc authService.AuthService) *Handler {
	return &Handler{Authuser: authSvc}
}
