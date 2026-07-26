package user_impl

import (
	"chatapp-backend/models"
	"context"
)

func (s *UserServiceImpl) ListUsers(ctx context.Context, currentUserID string) ([]models.UserInfo, error) {
	return s.repo.ListUsers(ctx, currentUserID)
}
