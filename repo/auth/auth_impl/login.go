package auth_impl

import (
	apperror "chatapp-backend/error"
	"chatapp-backend/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *AuthImpl) Login(ctx context.Context, email string) (*models.Users, error) {
	var user models.Users

	result := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}
