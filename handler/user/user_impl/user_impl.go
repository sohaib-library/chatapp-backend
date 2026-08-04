package user_impl

import (
	userHandler "github.com/sohaib-library/chatapp-backend/handler/user"
	userService "github.com/sohaib-library/chatapp-backend/service/user"
)

var _ userHandler.UserHandler = (*Handler)(nil)

type Handler struct {
	User userService.UserService
}

func NewHandler(svc userService.UserService) *Handler {
	return &Handler{User: svc}
}
