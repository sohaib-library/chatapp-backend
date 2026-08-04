package user

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

type UserService interface {
	ListUsers(ctx context.Context, currentUserID string) ([]models.UserInfo, error)
}
