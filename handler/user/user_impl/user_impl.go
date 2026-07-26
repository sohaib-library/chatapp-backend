package user_impl

import userService "chatapp-backend/service/user"

type Handler struct {
	User userService.UserService
}

func NewHandler(svc userService.UserService) *Handler {
	return &Handler{User: svc}
}
