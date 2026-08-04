package user_impl

import (
	"context"
	"errors"

	"github.com/sohaib-library/chatapp-backend/models"
	"gorm.io/gorm"
)

func (r *UserImpl) GetMe(ctx context.Context, userID string) (*models.UserInfo, error) {
	var user models.UserDB

	result := r.db.WithContext(ctx).
		Where("id = ?", userID).
		First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &models.UserInfo{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
