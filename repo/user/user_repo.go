package user

import (
	"chatapp-backend/models"
	"context"
)

type UserRepo interface {
	ListUsers(ctx context.Context, excludeUserID string) ([]models.UserInfo, error)
}
