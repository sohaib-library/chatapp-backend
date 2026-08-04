package auth_impl

import (
	"context"
	"errors"

	apperror "github.com/sohaib-library/chatapp-backend/error"
	"github.com/sohaib-library/chatapp-backend/models"

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
