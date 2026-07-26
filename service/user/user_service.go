package user

import (
	"chatapp-backend/models"
	"context"
)

type UserService interface {
	ListUsers(ctx context.Context, currentUserID string) ([]models.UserInfo, error)
}
