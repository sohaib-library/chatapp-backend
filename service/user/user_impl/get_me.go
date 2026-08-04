package user_impl

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

func (s *UserServiceImpl) GetMe(ctx context.Context, userID string) (*models.UserInfo, error) {
	return s.repo.GetMe(ctx, userID)
}
