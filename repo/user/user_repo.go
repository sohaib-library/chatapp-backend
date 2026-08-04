
package user

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

type UserRepo interface {
	ListUsers(ctx context.Context, excludeUserID string) ([]models.UserInfo, error)
	GetMe(ctx context.Context, userID string) (*models.UserInfo, error)
}
