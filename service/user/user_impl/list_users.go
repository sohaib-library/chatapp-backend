package user_impl

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

func (s *UserServiceImpl) ListUsers(ctx context.Context, currentUserID string) ([]models.UserInfo, error) {
	return s.repo.ListUsers(ctx, currentUserID)
}
