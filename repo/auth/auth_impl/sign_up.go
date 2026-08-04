package auth_impl

import (
	"context"
	"errors"

	apperror "github.com/sohaib-library/chatapp-backend/error"
	"github.com/sohaib-library/chatapp-backend/models"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AuthImpl) CreateUser(ctx context.Context, user models.Users) error {
	result := r.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) && pgErr.Code == "23505" {
			return apperror.ErrUserExists
		}
		return result.Error
	}
	return nil
}

func (r *AuthImpl) UserExists(ctx context.Context, email string) (bool, error) {
	var count int64
	result := r.db.WithContext(ctx).
		Model(&models.Users{}).
		Where("email = ?", email).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
